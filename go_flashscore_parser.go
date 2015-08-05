package go_flashscore_parser

import (
	"fmt"
	"log"
	"github.com/fedesog/webdriver"
	"time"
	"strings"
)

type matchBlock struct {
	countryName string
	leagueName string
	matches []match
}

type match struct {
	homeTeamName string
	awayTeamName string
	goalsHome string
	goalsAway string
	startTime string
}

const (
	FLASHSCORE_URL = "http://www.flashscore.ru"
	FLASHSCORE_LOAD_TIME = 10
)

func GetGames(chromeDriverPath string) []matchBlock {
	
	//adding chromedriver and opening page
	chromeDriver := webdriver.NewChromeDriver(chromeDriverPath)
	err := chromeDriver.Start()
	if err != nil {
	    log.Println(err)
	}
	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
	    log.Println(err)
	}
	err = session.Url(FLASHSCORE_URL)
	if err != nil {
	    log.Println(err)
	}
	
	//wait for page to load
	time.Sleep(FLASHSCORE_LOAD_TIME * time.Second)
	
	//array game blocks
	var blocks []matchBlock
	
	//find all blocks with games
	HTMLblocks, _ := session.FindElements(webdriver.CSS_Selector, "table.soccer")
	for i := 0; i < len(HTMLblocks); i++ {
		
		var matches []match
		
		//finding country name
		element, err := HTMLblocks[i].FindElement(webdriver.ClassName, "country_part")
		if err != nil {
			fmt.Println("Can not find country")
		}
		//deleting ":" from country name
		textValue, err := element.Text()
		if err != nil {
			fmt.Println("Can not extract text")
		}
		countryName := strings.Replace(textValue, ":", "", -1)
		//finding league name
		element, err = HTMLblocks[i].FindElement(webdriver.ClassName, "tournament_part")
		leagueName, err := element.Text()
		if err != nil {
			fmt.Println("Can not extract league name text")
		}
		
		//finding all games in current league
		matchTRs, err := HTMLblocks[i].FindElements(webdriver.CSS_Selector, "tbody tr")
		if err != nil {
			fmt.Println("Can not find matches in league")
		}
		for j := 0; j < len(matchTRs); j++ {
			//finding team names
			homeTeamNameEl, err := matchTRs[j].FindElement(webdriver.CSS_Selector, "span.padr")
			if err != nil {
				fmt.Println("Can not find homeTeamName")
			}
			awayTeamNameEl, err := matchTRs[j].FindElement(webdriver.CSS_Selector, "span.padl")
			if err != nil {
				fmt.Println("Can not find awayTeamName")
			}
			homeTeamName, err := homeTeamNameEl.Text()
			if err != nil {
				fmt.Println("Can not extract homeTeamName text")
			}
			awayTeamName, err := awayTeamNameEl.Text()
			if err != nil {
				fmt.Println("Can not extract awayTeamName text")
			}
			//finding scores
			tdScoreEl, err := matchTRs[j].FindElement(webdriver.ClassName, "score")
			if err != nil {
				fmt.Println("Can not find score element")
			}
			tdScoreText, err := tdScoreEl.Text()
			if err != nil {
				fmt.Println("Can not extract text score")
			}
			//if game is playing or finished
			goalsHome := ""
			goalsAway := ""
			tdLength := len(tdScoreText)
			if(tdLength == 5){
				goalsHome = string(tdScoreText[0]);
				goalsAway = string(tdScoreText[tdLength - 1]);
			}
			//finding start time
			tdTimeEl, err := matchTRs[j].FindElement(webdriver.ClassName, "time")
			if err != nil {
				fmt.Println("Can not time element")
			}
			startTime, err := tdTimeEl.Text()
			if err != nil {
				fmt.Println("Can not extract text from time")
			}
			matches = append(matches, match{homeTeamName:homeTeamName, awayTeamName:awayTeamName, goalsHome:goalsHome, goalsAway:goalsAway, startTime:startTime})
		} 
		
		blocks = append(blocks, matchBlock{countryName:countryName, leagueName: leagueName, matches: matches})
		
	}
	
	session.Delete()
	chromeDriver.Stop()
	
	return blocks;
	
}

func Show(blocks []matchBlock){

	for i:=0; i < len(blocks); i++ {
		fmt.Printf("%s %s\n", blocks[i].countryName, blocks[i].leagueName)
		for j:=0; j < len(blocks[i].matches); j++ {
			fmt.Printf("%s %s %s:%s %s\n", 
				blocks[i].matches[j].startTime, 
				blocks[i].matches[j].homeTeamName,
				blocks[i].matches[j].goalsHome,
				blocks[i].matches[j].goalsAway, 
				blocks[i].matches[j].awayTeamName)
		}
	}
	
}
