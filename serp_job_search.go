// Register at https://serpapi.com , login and get your API key 
package main

import (
  g "github.com/serpapi/google-search-results-golang"
  "fmt"
  //"github.com/antchfx/jsonquery"
  "encoding/json"
  "os"
)

func main(){
  q:=string("")

if len(os.Args[0:])<2 {
  fmt.Println("Run like:\n"+os.Args[0]+" erlang programmer tx\n")
 return 
}
 args:=os.Args[1:len(os.Args[0:])]
 for i:= range args {q=q+" "+args[i]
                    }
 fmt.Printf("query =%v\n",q)
  parameter := map[string]string{
    "engine": "google_jobs",
    "q": q, //"erlang dallas tx",//"programmer golang lewisville tx",
    "hl": "en",
    "api_key": "your api token here",
  }

  search := g.NewGoogleSearch(parameter, "your api token here")
  results, err := search.GetJSON()
//  jobs_results := results["jobs_results"].([]interface{})

  if (err==nil) {fmt.Printf("%v\n----------------------\n",results)
                 fmt.Printf("%#v\n----------------------\n",results)
                }
  if (err==nil) { r,err_1:= json.Marshal(results)
                 if (err_1==nil) {fmt.Printf("%v\n",string(r))}
                }

  if (err!=nil) { e,err_2:= json.Marshal(err)
                 if (err_2==nil) { fmt.Printf("%v",string(e))}
                }

}