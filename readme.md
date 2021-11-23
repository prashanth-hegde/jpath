# jpath

With modern software, json has become a de-facto standard 
for transport between multiple systems. 
From UI including backend miccroservices.
But analyzing json data is painful. Which is why there are so many competing standards
vying for attention, including yaml, toml etc. 
<br>
But parsing json didn't have to be this hard. Back in the days of xml, we had xpath, 
which made searching and parsing xml documents easier. For json, there is very little.

There is `jq` and more recently `jsonpath` 
<br>
`jq` is feature rich owing to active development
<br>
`jpath` is an attempt at providing a very easy to use DSL and parser for json documents

## Concepts

jpath is optimized for find/filtering json documents. And ease-of-use of such operations. 
jpath is inspired by the [jsonpath spec](https://github.com/json-path/JsonPath), 
and tries to improve upon it by reducing/eliminating a lot of special characters (`$`, `"`, `@` etc.)
needed to get to the desired output 
<br>
An attempt is made to make the user expression as simple as possible without having to worry about arrays and objects. 
So, as a foundational tenet, all json documents are treated as arrays. If the input json doc is an object, 
it is internally represented as an array of one json document. Likewise, the output is always wrapped into an array
to make it easy to use it in further piped expressions.

In order to make the user expressions simpler, we have made the following enhancements to the program
* very limited special characters to know and learn
* no need to check whether a document is an object or array
* no need to wrap strings in double quotes - simplifies the expression
* space characters in expressions are optional - for better readability
* inline operations on json - reduces copying of data and leads to faster processing

Note: [Input json is provided by this api](https://randomuser.me/api/?results=10)

### Concepts - syntax - [json document](https://randomuser.me/api/?results=10)

|Name        | Operator | Example                        | Description   |
|------------|----------|--------------------------------|---------------|
| Separator  | `.`      | `results.name`                 | descends to the given operators while maintaining json array output format
| Filter*    |`[filter]`| `results[gender=female]`       | `[]` is an array, any expression inside the array forms the filter, see [supported filters](#concepts---supported-filters)
| Count      | `#`      | `results.#`                    | count of number of items in the array
| Slice      | `[1:3]`  | `results[1:2]`                 | sub-array from the given indexes
| Selection (TBD)| `{}` | `results.name.{first,last}`    | prints only the selected fields from the results

### Concepts - supported filters
|Name        | Operator | Example                        | Description   |
|------------|----------|--------------------------------|---------------|
| Equality   | `=`      | `results[name.title=Mr]`       | find results whose title is `Mr`
|Non-Equality| `!=`     | `results[name.title!=Mr]`      | find results whose title is _not_ `Mr`
| Less Than  | `<`      | `results[dob.age<60]`          | find results whose age is less than 60
|Greater Than| `>`      | `results[dob.age>20]`          | find results whose age is greater than 20 
| Regex Match| `~`      | `results[name.first~^P]`       | find results whose first name starts with a `P`

### Concepts - comparison with other tools
| Description                         | jq                      | jsonpath                          | jpath    |
|-------------------------------------|-------------------------|-----------------------------------|-----------|
| find females in the result          | `.results[] \| select(.gender == "female")` | `$.results[?(@.gender=="female")]` | `results[gender=female]`
| find count of results from above    | N/A                     | N/A                               | `results[gender=female].#`
| find name, dob and cell# for Males  | `.results[] \| select(.name.title = "Mr") \| {first:.name.first,last:.name.last,dob:.dob.date,cell:.cell}` | N/A | `'results[name.title=Mr].{name.first,name.last,dob.date,cell}'`

TODO: more documentation pending

### Credits
Credits and heartfelt thanks to the following opensource projects that makes `jpath` a reality

1. [colorjson](https://github.com/TylerBrock/colorjson) - colored terminal output
2. [tablewriter](https://github.com/olekukonko/tablewriter) - table output
3. [cobra](https://github.com/spf13/cobra) - cli parsing
