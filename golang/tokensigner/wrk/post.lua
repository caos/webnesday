-- example HTTP POST script which demonstrates setting the
-- HTTP method, body, and adding a header

wrk.method = "POST"
wrk.body   = '{"text": "hello webnesday","signed": "3291498fbe0b7b92b941c11d69b6810fa536abd157c0c358b7c3658c7c8b147f9cf9b41ae2177082774593493062b7307865ab674a6b11a71d8df5f8703069a18d269ec34489dfb0f8e009489e34fa9a522f791bbf616d4d30d5eabaf0450704af62b45694514bb90d3039814df46d95bf03794b7a035aaa6145dbea858ca3cf"}'
wrk.headers["Content-Type"] = "application/json"

