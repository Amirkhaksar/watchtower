package main

// --oos (remove out of scope assets)
var flagHelp = `Flags:
   --body "bodystring" (request body)
   --body-file "filename" (request body file name)
   --public-target "program_name" (add public target by name)
   --method string (http request method)
		
   --compare "filename" (compare response)
   --rc (reverse compare)
		  
   --limit (disable default loop)
   --loop (get all pages)
		
   --count (show count of results)
   --json (show output as json)
   --cdn (add cdn=ture)
   --help (help of flags)
   --total (add total=true)
   --internal (add internal=true)
   --no-limit (add not_limit=true)
		
   --date string (set date of results)
   --provider string (filter by providers)
   --status string (filter by status)
   --title string (filter by title)
   --technology (filter by technology)
   --response-headers (filter by response header, format: HEADER1:value; HEADER2:value2)
   --content-type (filter by content-type header)
   --server (filter by server)

   --exclude-domain string (exclude a domain from results)
   --exclude-provider string (exclude a provider from results)
   --exclude-scope string (exclude a scope from results)`

var getHelp = `watchtower get single target {{target_name}}
watchtower get single subdomain {{subdomain}}
watchtower get single live {{domain}}
watchtower get single http {{subdomain}}

watchtower get subdomains domain {{domain}}
watchtower get subdomains scope {{scope}}
watchtower get subdomains all

watchtower get lives scope {{scope}}
watchtower get lives domain {{domain}}
watchtower get lives all

watchtower get http scope {{scope}}
watchtower get http domain {{domain}}
watchtower get http all

watchtower get latest subdomains domain {{domain}}
watchtower get latest subdomains scope {{scope}}
watchtower get latest subdomains all

watchtower get latest lives domain {{domain}}
watchtower get latest lives scope {{scope}}
watchtower get latest lives all

watchtower get latest http domain {{domain}}
watchtower get latest http scope {{scope}}
watchtower get latest http all

watchtower get targets list
watchtower get targets public all
watchtower get targets public platform {{platform}}

watchtower get fresh subdomains all 
watchtower get fresh subdomains scope {{scope}} 
watchtower get fresh subdomains domain {{domain}} 

watchtower get fresh lives all 
watchtower get fresh lives scope {{scope}} 
watchtower get fresh lives domain {{domain}} 

watchtower get fresh http all 
watchtower get fresh http scope {{scope}} 
watchtower get fresh http domain {{domain}} 

watchtower get statistics sqs
watchtower get technologies list
`

var regexHelp = `watchtower regexp list
watchtower regexp apply -body-file body.txt
watchtower regexp test scope {{scope}} -body-file body.txt
watchtower regexp test all  -body-file body.txt
`

var orchHelp = `watchtower orch clean database scope {{scope}}
watchtower orch clean database all 
watchtower orch clean tags scope {{scope}}
watchtower orch clean tags all
watchtower orch clean ips scope {{scope}}
watchtower orch clean ips all
watchtower orch clean domains scope {{scope}}
watchtower orch clean domains all

watchtower orch passive enum all 
watchtower orch passive enum scope {{scope}}

watchtower orch resolution all
watchtower orch resolution scope {{scope}}

watchtower orch http all
watchtower orch http scope {{scope}}

watchtower orch targets update
watchtower orch push subenum
`

var putHelp = `watchtower put resolution 
watchtower put subdomain {{scope}}
watchtower put orch resolution
watchtower put orch http
`
var targetHelp = `watchtower target delete {{target_name}}
watchtower target create --body-file target.txt`
