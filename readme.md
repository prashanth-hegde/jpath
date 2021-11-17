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
`jq` is feature rich owing to active development 
`jpath` is an attempt at providing a very easy to use DSL and parser standard for json documents


## Concepts

jpath is optimized for find/filtering json documents. And ease-of-use of such operations. 
jpath is inspired by the [jsonpath spec](https://github.com/json-path/JsonPath), 
and tries to improve upon it by reducing/eliminating a lot of special characters
needed to get to the desired document 
<br>
An attempt is made to make the user expression as simple as possible without having to worry about arrays and objects. 
So, as a foundational tenet, all json documents are treated as arrays. If the input json doc is an object, 
it is internally represented as an array of one json document. Likewise, the output is always wrapped into an array
to make it easy to use it in further piped expressions.

In order to make the user expressions simpler, we have made the following enhancements to the program
* very limited special characters to know and learn
* no need to check whether a document is an object or array
* no need to wrap strings in double quotes - simplifies the expression
* space characters in expressions are optional

Note: [Input json is provided by this api](https://randomuser.me/api/?results=10)

### Concepts - syntax - [json document](https://randomuser.me/api/?results=10)

|Name        | Operator | Example                        | Description   |
|------------|----------|--------------------------------|---------------|
| Separator  | `.`      | `results.name`                 | descends to the given operators while maintaining json array output format
| Filter*    |`[filter]`| `results[gender=female]`       | `[]` is an array, any expression inside the array forms the filter
| Count      | `#`      | `results.#`                    | count of number of items in the array
| Composition (TBD)| `{}`     | `results.name.{first:.first}`  | composes a new json structure from the given param

### Concepts - supported operations
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
| find females in the result          | .results[] &#124; select(.gender == "female") | `$.results[?(@.gender=="female")]` | `results[gender=female]`
| find count of results from above    | N/A                     | N/A                               | `results[gender=female].#`
| todo: more examples

TODO: more documentation pending