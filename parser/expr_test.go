package parser

import (
	"fmt"
	"jpath/common"
	"reflect"
	"strings"
	"testing"
)

func TestParseExpression(t *testing.T) {
	testData := []struct {
		name     string
		input    string
		expected string
	}{
		{"base", `results`, `results`},
		{"nested", `results.name`, `results|name`},
		{"padding", `.results.name`, `results|name`},
		{"filter", `results[gender=female]`, `results[gender=female]`},
		{"nestedfilter", `..results[name.title=Miss]`, `results[name.title=Miss]`},
	}
	for _, testcase := range testData {
		output, _ := parseExpression(testcase.input)
		expectedTokens := strings.Split(testcase.expected, "|")
		fmt.Println(output)
		if !reflect.DeepEqual(expectedTokens, output) {
			t.Errorf("%s --> failed", testcase.name)
		}
	}
}

func TestProcessExpression(t *testing.T) {
	testJson, _ := common.Tokenize([]byte(`{"results":[{"gender":"female","name":{"title":"Mrs","first":"Melodie","last":"Gagné"},"location":{"street":{"number":4594,"name":"St. Lawrence Ave"},"city":"Killarney","state":"Ontario","country":"Canada","postcode":"L5C 7L2","coordinates":{"latitude":"-63.9167","longitude":"110.8605"},"timezone":{"offset":"+5:45","description":"Kathmandu"}},"email":"melodie.gagne@example.com","login":{"uuid":"d56c7d7a-b4d4-4250-9d28-0738ac3a1b77","username":"orangebear509","password":"bigboy","salt":"r2rkwcwY","md5":"bea5a3bffe48b1e69c817688e929cf98","sha1":"9e67c5ddd7a1cdb28866f4ef40d0be64f389a555","sha256":"3a83320997a8bf8c860c6f2b813cc7c2dbaa68a958082109d0ef973b62eedc64"},"dob":{"date":"1956-12-07T12:09:05.427Z","age":65},"registered":{"date":"2011-10-17T13:35:21.088Z","age":10},"phone":"726-074-4524","cell":"779-777-0657","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/women/62.jpg","medium":"https://randomuser.me/api/portraits/med/women/62.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/women/62.jpg"},"nat":"CA"},{"gender":"male","name":{"title":"Mr","first":"Hardy","last":"Sauter"},"location":{"street":{"number":3759,"name":"Raiffeisenstraße"},"city":"Schwandorf","state":"Hamburg","country":"Germany","postcode":43773,"coordinates":{"latitude":"9.3017","longitude":"-78.3350"},"timezone":{"offset":"+5:00","description":"Ekaterinburg, Islamabad, Karachi, Tashkent"}},"email":"hardy.sauter@example.com","login":{"uuid":"7bbfbcb0-a9c7-491b-b347-6bb6cc560903","username":"bluebear598","password":"freeuser","salt":"aMcmc5p1","md5":"321930e49c1e0797e27f97c5cffbb744","sha1":"f00f7f2f40a1b63c799d849dc7b1ac7a09124931","sha256":"8eaf321a3b77fe78de3124ce0d54969d14dc3e327e74801d826495baf6423b00"},"dob":{"date":"1955-09-06T14:26:49.215Z","age":66},"registered":{"date":"2014-02-27T05:56:03.362Z","age":7},"phone":"0461-2742163","cell":"0173-7395805","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/men/16.jpg","medium":"https://randomuser.me/api/portraits/med/men/16.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/16.jpg"},"nat":"DE"},{"gender":"male","name":{"title":"Mr","first":"Porfírio","last":"Pereira"},"location":{"street":{"number":6062,"name":"Travessa dos Açorianos"},"city":"Nossa Senhora do Socorro","state":"Pernambuco","country":"Brazil","postcode":17855,"coordinates":{"latitude":"-81.9210","longitude":"150.2790"},"timezone":{"offset":"+6:00","description":"Almaty, Dhaka, Colombo"}},"email":"porfirio.pereira@example.com","login":{"uuid":"4d52083a-2692-43b7-83a5-5c5e9d0815cb","username":"orangelion732","password":"frosty","salt":"qnsQL5c0","md5":"a96f887a42ee546badb3f43681c998e9","sha1":"fc6528f28d113e7ba796f06809994ff15fe077d2","sha256":"7b886f2472a6e4d54b5b5037bf3fe83780465eeb2b305d614f295daf91724cb5"},"dob":{"date":"1980-09-06T08:26:56.458Z","age":41},"registered":{"date":"2004-09-02T17:44:40.628Z","age":17},"phone":"(93) 1456-9695","cell":"(31) 0998-5800","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/men/87.jpg","medium":"https://randomuser.me/api/portraits/med/men/87.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/87.jpg"},"nat":"BR"},{"gender":"male","name":{"title":"Mr","first":"Loïs","last":"Blanc"},"location":{"street":{"number":9030,"name":"Rue de L'Abbé-De-L'Épée"},"city":"Asnières-sur-Seine","state":"Finistère","country":"France","postcode":89952,"coordinates":{"latitude":"50.9744","longitude":"38.2468"},"timezone":{"offset":"+9:00","description":"Tokyo, Seoul, Osaka, Sapporo, Yakutsk"}},"email":"lois.blanc@example.com","login":{"uuid":"f991434b-9e0f-4588-ba2e-f6179af9e842","username":"brownkoala788","password":"tigger","salt":"fxGzgqe3","md5":"8ba64f770c8b99c9eba1528d76c5888a","sha1":"90334001cdd9f5f021e3832969885416328da39c","sha256":"10e916d5521a6db13318c5f7113a6a89160fcd08c687e2eec2219785c3bcb704"},"dob":{"date":"1969-01-13T18:35:55.654Z","age":52},"registered":{"date":"2017-02-07T16:01:02.108Z","age":4},"phone":"04-76-20-47-49","cell":"06-13-55-64-94","id":{"name":"INSEE","value":"1NNaN46527313 64"},"picture":{"large":"https://randomuser.me/api/portraits/men/90.jpg","medium":"https://randomuser.me/api/portraits/med/men/90.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/90.jpg"},"nat":"FR"},{"gender":"female","name":{"title":"Mrs","first":"Scarlett","last":"Ramos"},"location":{"street":{"number":6288,"name":"E Pecan St"},"city":"Queanbeyan","state":"Queensland","country":"Australia","postcode":8613,"coordinates":{"latitude":"-84.4282","longitude":"162.1111"},"timezone":{"offset":"-6:00","description":"Central Time (US & Canada), Mexico City"}},"email":"scarlett.ramos@example.com","login":{"uuid":"ed8d0175-81ee-4814-a168-6141ee47c7de","username":"brownsnake968","password":"mission","salt":"aMOaO9fR","md5":"2e9ba6b0c914923c170bd405761f31ad","sha1":"29068dea51109fec794defcfa81481ba004a8e17","sha256":"9018f4f9d3b37a6a4211174936d36213368ebba68cb1be7aae04decf8991e822"},"dob":{"date":"1987-05-02T15:04:52.481Z","age":34},"registered":{"date":"2003-07-23T08:27:48.875Z","age":18},"phone":"02-3890-8160","cell":"0491-054-120","id":{"name":"TFN","value":"045760989"},"picture":{"large":"https://randomuser.me/api/portraits/women/63.jpg","medium":"https://randomuser.me/api/portraits/med/women/63.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/women/63.jpg"},"nat":"AU"},{"gender":"male","name":{"title":"Mr","first":"Glen","last":"Howard"},"location":{"street":{"number":1180,"name":"George Street"},"city":"Swords","state":"Wexford","country":"Ireland","postcode":99620,"coordinates":{"latitude":"29.7085","longitude":"75.1163"},"timezone":{"offset":"+5:30","description":"Bombay, Calcutta, Madras, New Delhi"}},"email":"glen.howard@example.com","login":{"uuid":"a316c6e2-588b-46f7-bb84-b5259809db63","username":"greenbird880","password":"tttttt","salt":"gXFbTT9E","md5":"740e62a443776e531c32e8a27936cd32","sha1":"8cae36827b4626ee5d4bb33ec9a950b442c77be0","sha256":"299b67ef37fb0dfaee5069bf42cc4760cf4683b3f747ca5df974c2e22ef5940a"},"dob":{"date":"1990-02-13T12:02:30.019Z","age":31},"registered":{"date":"2009-06-27T06:37:56.602Z","age":12},"phone":"061-787-3251","cell":"081-991-2405","id":{"name":"PPS","value":"5290365T"},"picture":{"large":"https://randomuser.me/api/portraits/men/49.jpg","medium":"https://randomuser.me/api/portraits/med/men/49.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/49.jpg"},"nat":"IE"},{"gender":"female","name":{"title":"Mrs","first":"Johanne","last":"Leonhardt"},"location":{"street":{"number":3194,"name":"Königsberger Straße"},"city":"Bad Dürrheim","state":"Bremen","country":"Germany","postcode":77926,"coordinates":{"latitude":"23.1901","longitude":"-101.7833"},"timezone":{"offset":"-5:00","description":"Eastern Time (US & Canada), Bogota, Lima"}},"email":"johanne.leonhardt@example.com","login":{"uuid":"b6709a0a-f1b9-43df-81a3-5df0c291e209","username":"crazytiger533","password":"portland","salt":"02lLcotQ","md5":"7a7ecfc38af3758df76249fa3dd2644d","sha1":"9ead20210f0148e30a7698e85202a6f6e9b937ca","sha256":"4817a46f630a5c53803e8deb28e58e9c6b3db8961640ae533982072b16f61c87"},"dob":{"date":"1964-08-09T08:11:15.101Z","age":57},"registered":{"date":"2004-01-17T03:23:18.800Z","age":17},"phone":"0678-4241451","cell":"0179-6896607","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/women/46.jpg","medium":"https://randomuser.me/api/portraits/med/women/46.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/women/46.jpg"},"nat":"DE"},{"gender":"male","name":{"title":"Mr","first":"Byron","last":"Cooper"},"location":{"street":{"number":7641,"name":"Lakeshore Rd"},"city":"St. Petersburg","state":"Colorado","country":"United States","postcode":85898,"coordinates":{"latitude":"-13.9778","longitude":"-137.0939"},"timezone":{"offset":"+8:00","description":"Beijing, Perth, Singapore, Hong Kong"}},"email":"byron.cooper@example.com","login":{"uuid":"40069ca0-fefd-4976-ab13-dff1c80bbf7e","username":"bigbird675","password":"nelson","salt":"h0ExMxtO","md5":"24dd7367f95c4377b2fef8e65cbd0c1b","sha1":"8c416504804eddc891da083eb3461733738fdc02","sha256":"96ce2b753601e405c00301242398c24551260208ac4799d78aacc2f8b87a7595"},"dob":{"date":"1992-02-22T10:15:55.954Z","age":29},"registered":{"date":"2007-03-13T17:08:10.687Z","age":14},"phone":"(230)-039-4521","cell":"(532)-586-9671","id":{"name":"SSN","value":"120-86-5372"},"picture":{"large":"https://randomuser.me/api/portraits/men/74.jpg","medium":"https://randomuser.me/api/portraits/med/men/74.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/74.jpg"},"nat":"US"},{"gender":"female","name":{"title":"Miss","first":"Yesim","last":"Nijenkamp"},"location":{"street":{"number":5373,"name":"1e Zandwijkje"},"city":"Zoutkamp","state":"Limburg","country":"Netherlands","postcode":80936,"coordinates":{"latitude":"-19.9554","longitude":"-162.6304"},"timezone":{"offset":"+3:30","description":"Tehran"}},"email":"yesim.nijenkamp@example.com","login":{"uuid":"dd67bceb-5392-4ee2-8708-94e04a84b887","username":"bigmouse583","password":"chaos1","salt":"wuRC9yFI","md5":"19a98b6b89a6aa21b94cf758f9cff7dc","sha1":"f7824391d9c48dd96bc8c3b1e2d30520c4999506","sha256":"36e88cd516175132231b1cd867bfecbd1a1b25b1cef1b2e85f38b0ce0a9ba4e0"},"dob":{"date":"1980-11-07T08:04:26.249Z","age":41},"registered":{"date":"2016-11-13T19:57:41.510Z","age":5},"phone":"(025)-795-8551","cell":"(144)-699-7915","id":{"name":"BSN","value":"88093887"},"picture":{"large":"https://randomuser.me/api/portraits/women/53.jpg","medium":"https://randomuser.me/api/portraits/med/women/53.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/women/53.jpg"},"nat":"NL"},{"gender":"male","name":{"title":"Mr","first":"Koray","last":"Adal"},"location":{"street":{"number":8407,"name":"Talak Göktepe Cd"},"city":"Tokat","state":"Bolu","country":"Turkey","postcode":34303,"coordinates":{"latitude":"-45.7986","longitude":"-65.0078"},"timezone":{"offset":"-4:00","description":"Atlantic Time (Canada), Caracas, La Paz"}},"email":"koray.adal@example.com","login":{"uuid":"804cfd39-034a-4b4e-b517-acd6373682ea","username":"heavyladybug985","password":"rhodes","salt":"S9IMfke6","md5":"f11651ff3ac9462bb9368b6f0e182ddb","sha1":"bc916144a87208e46539234eff95b750ffbbad0e","sha256":"a83c8c26a27793dab9e97c482a313bd755b0e799336861a3ccd45a6bf0e1d4cf"},"dob":{"date":"1983-09-02T03:40:15.660Z","age":38},"registered":{"date":"2005-01-20T11:36:41.891Z","age":16},"phone":"(130)-567-8057","cell":"(434)-340-1183","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/men/73.jpg","medium":"https://randomuser.me/api/portraits/med/men/73.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/73.jpg"},"nat":"TR"}],"info":{"seed":"17b3298ce93fa1a4","results":10,"page":1,"version":"1.3"}}`))
	testData := []struct {
		name     string
		input    string
		expected []string
	}{
		{"base", `info`, []string{`{"seed":"17b3298ce93fa1a4","results":10,"page":1,"version":"1.3"}`}},
		{"space", `info  `, []string{`{"seed":"17b3298ce93fa1a4","results":10,"page":1,"version":"1.3"}`}},
		{"nested", `results.gender`, []string{"female", "male", "male", "male", "female", "male", "female", "male", "female", "male"}},
		{"deepnested", `results.name.title`, []string{"Mrs", "Mr", "Mr", "Mr", "Mrs", "Mr", "Mrs", "Mr", "Miss", "Mr"}},
		{"nestedsub", `results.name[title=Mrs]`, []string{`{"title":"Mrs","first":"Melodie","last":"Gagné"}`, `{"title":"Mrs","first":"Scarlett","last":"Ramos"}`, `{"title":"Mrs","first":"Johanne","last":"Leonhardt"}`}},

		// select
		{"selectbase", `results[name.title=Mrs].{cell,dob}`, []string{`{"cell":"779-777-0657","dob":{"date":"1956-12-07T12:09:05.427Z","age":65}}`, `{"cell":"0491-054-120","dob":{"date":"1987-05-02T15:04:52.481Z","age":34}}`, `{"cell":"0179-6896607","dob":{"date":"1964-08-09T08:11:15.101Z","age":57}}`}},
		{"array", `{results,info}`, []string{`{"results":[{"gender":"female","name":{"title":"Mrs","first":"Melodie","last":"Gagné"},"location":{"street":{"number":4594,"name":"St. Lawrence Ave"},"city":"Killarney","state":"Ontario","country":"Canada","postcode":"L5C 7L2","coordinates":{"latitude":"-63.9167","longitude":"110.8605"},"timezone":{"offset":"+5:45","description":"Kathmandu"}},"email":"melodie.gagne@example.com","login":{"uuid":"d56c7d7a-b4d4-4250-9d28-0738ac3a1b77","username":"orangebear509","password":"bigboy","salt":"r2rkwcwY","md5":"bea5a3bffe48b1e69c817688e929cf98","sha1":"9e67c5ddd7a1cdb28866f4ef40d0be64f389a555","sha256":"3a83320997a8bf8c860c6f2b813cc7c2dbaa68a958082109d0ef973b62eedc64"},"dob":{"date":"1956-12-07T12:09:05.427Z","age":65},"registered":{"date":"2011-10-17T13:35:21.088Z","age":10},"phone":"726-074-4524","cell":"779-777-0657","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/women/62.jpg","medium":"https://randomuser.me/api/portraits/med/women/62.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/women/62.jpg"},"nat":"CA"},{"gender":"male","name":{"title":"Mr","first":"Hardy","last":"Sauter"},"location":{"street":{"number":3759,"name":"Raiffeisenstraße"},"city":"Schwandorf","state":"Hamburg","country":"Germany","postcode":43773,"coordinates":{"latitude":"9.3017","longitude":"-78.3350"},"timezone":{"offset":"+5:00","description":"Ekaterinburg, Islamabad, Karachi, Tashkent"}},"email":"hardy.sauter@example.com","login":{"uuid":"7bbfbcb0-a9c7-491b-b347-6bb6cc560903","username":"bluebear598","password":"freeuser","salt":"aMcmc5p1","md5":"321930e49c1e0797e27f97c5cffbb744","sha1":"f00f7f2f40a1b63c799d849dc7b1ac7a09124931","sha256":"8eaf321a3b77fe78de3124ce0d54969d14dc3e327e74801d826495baf6423b00"},"dob":{"date":"1955-09-06T14:26:49.215Z","age":66},"registered":{"date":"2014-02-27T05:56:03.362Z","age":7},"phone":"0461-2742163","cell":"0173-7395805","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/men/16.jpg","medium":"https://randomuser.me/api/portraits/med/men/16.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/16.jpg"},"nat":"DE"},{"gender":"male","name":{"title":"Mr","first":"Porfírio","last":"Pereira"},"location":{"street":{"number":6062,"name":"Travessa dos Açorianos"},"city":"Nossa Senhora do Socorro","state":"Pernambuco","country":"Brazil","postcode":17855,"coordinates":{"latitude":"-81.9210","longitude":"150.2790"},"timezone":{"offset":"+6:00","description":"Almaty, Dhaka, Colombo"}},"email":"porfirio.pereira@example.com","login":{"uuid":"4d52083a-2692-43b7-83a5-5c5e9d0815cb","username":"orangelion732","password":"frosty","salt":"qnsQL5c0","md5":"a96f887a42ee546badb3f43681c998e9","sha1":"fc6528f28d113e7ba796f06809994ff15fe077d2","sha256":"7b886f2472a6e4d54b5b5037bf3fe83780465eeb2b305d614f295daf91724cb5"},"dob":{"date":"1980-09-06T08:26:56.458Z","age":41},"registered":{"date":"2004-09-02T17:44:40.628Z","age":17},"phone":"(93) 1456-9695","cell":"(31) 0998-5800","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/men/87.jpg","medium":"https://randomuser.me/api/portraits/med/men/87.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/87.jpg"},"nat":"BR"},{"gender":"male","name":{"title":"Mr","first":"Loïs","last":"Blanc"},"location":{"street":{"number":9030,"name":"Rue de L'Abbé-De-L'Épée"},"city":"Asnières-sur-Seine","state":"Finistère","country":"France","postcode":89952,"coordinates":{"latitude":"50.9744","longitude":"38.2468"},"timezone":{"offset":"+9:00","description":"Tokyo, Seoul, Osaka, Sapporo, Yakutsk"}},"email":"lois.blanc@example.com","login":{"uuid":"f991434b-9e0f-4588-ba2e-f6179af9e842","username":"brownkoala788","password":"tigger","salt":"fxGzgqe3","md5":"8ba64f770c8b99c9eba1528d76c5888a","sha1":"90334001cdd9f5f021e3832969885416328da39c","sha256":"10e916d5521a6db13318c5f7113a6a89160fcd08c687e2eec2219785c3bcb704"},"dob":{"date":"1969-01-13T18:35:55.654Z","age":52},"registered":{"date":"2017-02-07T16:01:02.108Z","age":4},"phone":"04-76-20-47-49","cell":"06-13-55-64-94","id":{"name":"INSEE","value":"1NNaN46527313 64"},"picture":{"large":"https://randomuser.me/api/portraits/men/90.jpg","medium":"https://randomuser.me/api/portraits/med/men/90.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/90.jpg"},"nat":"FR"},{"gender":"female","name":{"title":"Mrs","first":"Scarlett","last":"Ramos"},"location":{"street":{"number":6288,"name":"E Pecan St"},"city":"Queanbeyan","state":"Queensland","country":"Australia","postcode":8613,"coordinates":{"latitude":"-84.4282","longitude":"162.1111"},"timezone":{"offset":"-6:00","description":"Central Time (US & Canada), Mexico City"}},"email":"scarlett.ramos@example.com","login":{"uuid":"ed8d0175-81ee-4814-a168-6141ee47c7de","username":"brownsnake968","password":"mission","salt":"aMOaO9fR","md5":"2e9ba6b0c914923c170bd405761f31ad","sha1":"29068dea51109fec794defcfa81481ba004a8e17","sha256":"9018f4f9d3b37a6a4211174936d36213368ebba68cb1be7aae04decf8991e822"},"dob":{"date":"1987-05-02T15:04:52.481Z","age":34},"registered":{"date":"2003-07-23T08:27:48.875Z","age":18},"phone":"02-3890-8160","cell":"0491-054-120","id":{"name":"TFN","value":"045760989"},"picture":{"large":"https://randomuser.me/api/portraits/women/63.jpg","medium":"https://randomuser.me/api/portraits/med/women/63.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/women/63.jpg"},"nat":"AU"},{"gender":"male","name":{"title":"Mr","first":"Glen","last":"Howard"},"location":{"street":{"number":1180,"name":"George Street"},"city":"Swords","state":"Wexford","country":"Ireland","postcode":99620,"coordinates":{"latitude":"29.7085","longitude":"75.1163"},"timezone":{"offset":"+5:30","description":"Bombay, Calcutta, Madras, New Delhi"}},"email":"glen.howard@example.com","login":{"uuid":"a316c6e2-588b-46f7-bb84-b5259809db63","username":"greenbird880","password":"tttttt","salt":"gXFbTT9E","md5":"740e62a443776e531c32e8a27936cd32","sha1":"8cae36827b4626ee5d4bb33ec9a950b442c77be0","sha256":"299b67ef37fb0dfaee5069bf42cc4760cf4683b3f747ca5df974c2e22ef5940a"},"dob":{"date":"1990-02-13T12:02:30.019Z","age":31},"registered":{"date":"2009-06-27T06:37:56.602Z","age":12},"phone":"061-787-3251","cell":"081-991-2405","id":{"name":"PPS","value":"5290365T"},"picture":{"large":"https://randomuser.me/api/portraits/men/49.jpg","medium":"https://randomuser.me/api/portraits/med/men/49.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/49.jpg"},"nat":"IE"},{"gender":"female","name":{"title":"Mrs","first":"Johanne","last":"Leonhardt"},"location":{"street":{"number":3194,"name":"Königsberger Straße"},"city":"Bad Dürrheim","state":"Bremen","country":"Germany","postcode":77926,"coordinates":{"latitude":"23.1901","longitude":"-101.7833"},"timezone":{"offset":"-5:00","description":"Eastern Time (US & Canada), Bogota, Lima"}},"email":"johanne.leonhardt@example.com","login":{"uuid":"b6709a0a-f1b9-43df-81a3-5df0c291e209","username":"crazytiger533","password":"portland","salt":"02lLcotQ","md5":"7a7ecfc38af3758df76249fa3dd2644d","sha1":"9ead20210f0148e30a7698e85202a6f6e9b937ca","sha256":"4817a46f630a5c53803e8deb28e58e9c6b3db8961640ae533982072b16f61c87"},"dob":{"date":"1964-08-09T08:11:15.101Z","age":57},"registered":{"date":"2004-01-17T03:23:18.800Z","age":17},"phone":"0678-4241451","cell":"0179-6896607","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/women/46.jpg","medium":"https://randomuser.me/api/portraits/med/women/46.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/women/46.jpg"},"nat":"DE"},{"gender":"male","name":{"title":"Mr","first":"Byron","last":"Cooper"},"location":{"street":{"number":7641,"name":"Lakeshore Rd"},"city":"St. Petersburg","state":"Colorado","country":"United States","postcode":85898,"coordinates":{"latitude":"-13.9778","longitude":"-137.0939"},"timezone":{"offset":"+8:00","description":"Beijing, Perth, Singapore, Hong Kong"}},"email":"byron.cooper@example.com","login":{"uuid":"40069ca0-fefd-4976-ab13-dff1c80bbf7e","username":"bigbird675","password":"nelson","salt":"h0ExMxtO","md5":"24dd7367f95c4377b2fef8e65cbd0c1b","sha1":"8c416504804eddc891da083eb3461733738fdc02","sha256":"96ce2b753601e405c00301242398c24551260208ac4799d78aacc2f8b87a7595"},"dob":{"date":"1992-02-22T10:15:55.954Z","age":29},"registered":{"date":"2007-03-13T17:08:10.687Z","age":14},"phone":"(230)-039-4521","cell":"(532)-586-9671","id":{"name":"SSN","value":"120-86-5372"},"picture":{"large":"https://randomuser.me/api/portraits/men/74.jpg","medium":"https://randomuser.me/api/portraits/med/men/74.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/74.jpg"},"nat":"US"},{"gender":"female","name":{"title":"Miss","first":"Yesim","last":"Nijenkamp"},"location":{"street":{"number":5373,"name":"1e Zandwijkje"},"city":"Zoutkamp","state":"Limburg","country":"Netherlands","postcode":80936,"coordinates":{"latitude":"-19.9554","longitude":"-162.6304"},"timezone":{"offset":"+3:30","description":"Tehran"}},"email":"yesim.nijenkamp@example.com","login":{"uuid":"dd67bceb-5392-4ee2-8708-94e04a84b887","username":"bigmouse583","password":"chaos1","salt":"wuRC9yFI","md5":"19a98b6b89a6aa21b94cf758f9cff7dc","sha1":"f7824391d9c48dd96bc8c3b1e2d30520c4999506","sha256":"36e88cd516175132231b1cd867bfecbd1a1b25b1cef1b2e85f38b0ce0a9ba4e0"},"dob":{"date":"1980-11-07T08:04:26.249Z","age":41},"registered":{"date":"2016-11-13T19:57:41.510Z","age":5},"phone":"(025)-795-8551","cell":"(144)-699-7915","id":{"name":"BSN","value":"88093887"},"picture":{"large":"https://randomuser.me/api/portraits/women/53.jpg","medium":"https://randomuser.me/api/portraits/med/women/53.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/women/53.jpg"},"nat":"NL"},{"gender":"male","name":{"title":"Mr","first":"Koray","last":"Adal"},"location":{"street":{"number":8407,"name":"Talak Göktepe Cd"},"city":"Tokat","state":"Bolu","country":"Turkey","postcode":34303,"coordinates":{"latitude":"-45.7986","longitude":"-65.0078"},"timezone":{"offset":"-4:00","description":"Atlantic Time (Canada), Caracas, La Paz"}},"email":"koray.adal@example.com","login":{"uuid":"804cfd39-034a-4b4e-b517-acd6373682ea","username":"heavyladybug985","password":"rhodes","salt":"S9IMfke6","md5":"f11651ff3ac9462bb9368b6f0e182ddb","sha1":"bc916144a87208e46539234eff95b750ffbbad0e","sha256":"a83c8c26a27793dab9e97c482a313bd755b0e799336861a3ccd45a6bf0e1d4cf"},"dob":{"date":"1983-09-02T03:40:15.660Z","age":38},"registered":{"date":"2005-01-20T11:36:41.891Z","age":16},"phone":"(130)-567-8057","cell":"(434)-340-1183","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/men/73.jpg","medium":"https://randomuser.me/api/portraits/med/men/73.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/73.jpg"},"nat":"TR"}],"info":{"seed":"17b3298ce93fa1a4","results":10,"page":1,"version":"1.3"}}`}},

		// compare
		{"gt", `results[dob.age>65]`, []string{`{"gender":"male","name":{"title":"Mr","first":"Hardy","last":"Sauter"},"location":{"street":{"number":3759,"name":"Raiffeisenstraße"},"city":"Schwandorf","state":"Hamburg","country":"Germany","postcode":43773,"coordinates":{"latitude":"9.3017","longitude":"-78.3350"},"timezone":{"offset":"+5:00","description":"Ekaterinburg, Islamabad, Karachi, Tashkent"}},"email":"hardy.sauter@example.com","login":{"uuid":"7bbfbcb0-a9c7-491b-b347-6bb6cc560903","username":"bluebear598","password":"freeuser","salt":"aMcmc5p1","md5":"321930e49c1e0797e27f97c5cffbb744","sha1":"f00f7f2f40a1b63c799d849dc7b1ac7a09124931","sha256":"8eaf321a3b77fe78de3124ce0d54969d14dc3e327e74801d826495baf6423b00"},"dob":{"date":"1955-09-06T14:26:49.215Z","age":66},"registered":{"date":"2014-02-27T05:56:03.362Z","age":7},"phone":"0461-2742163","cell":"0173-7395805","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/men/16.jpg","medium":"https://randomuser.me/api/portraits/med/men/16.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/16.jpg"},"nat":"DE"}`}},
		{"gte", `results[dob.age>=65]`, []string{`{"gender":"female","name":{"title":"Mrs","first":"Melodie","last":"Gagné"},"location":{"street":{"number":4594,"name":"St. Lawrence Ave"},"city":"Killarney","state":"Ontario","country":"Canada","postcode":"L5C 7L2","coordinates":{"latitude":"-63.9167","longitude":"110.8605"},"timezone":{"offset":"+5:45","description":"Kathmandu"}},"email":"melodie.gagne@example.com","login":{"uuid":"d56c7d7a-b4d4-4250-9d28-0738ac3a1b77","username":"orangebear509","password":"bigboy","salt":"r2rkwcwY","md5":"bea5a3bffe48b1e69c817688e929cf98","sha1":"9e67c5ddd7a1cdb28866f4ef40d0be64f389a555","sha256":"3a83320997a8bf8c860c6f2b813cc7c2dbaa68a958082109d0ef973b62eedc64"},"dob":{"date":"1956-12-07T12:09:05.427Z","age":65},"registered":{"date":"2011-10-17T13:35:21.088Z","age":10},"phone":"726-074-4524","cell":"779-777-0657","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/women/62.jpg","medium":"https://randomuser.me/api/portraits/med/women/62.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/women/62.jpg"},"nat":"CA"}`, `{"gender":"male","name":{"title":"Mr","first":"Hardy","last":"Sauter"},"location":{"street":{"number":3759,"name":"Raiffeisenstraße"},"city":"Schwandorf","state":"Hamburg","country":"Germany","postcode":43773,"coordinates":{"latitude":"9.3017","longitude":"-78.3350"},"timezone":{"offset":"+5:00","description":"Ekaterinburg, Islamabad, Karachi, Tashkent"}},"email":"hardy.sauter@example.com","login":{"uuid":"7bbfbcb0-a9c7-491b-b347-6bb6cc560903","username":"bluebear598","password":"freeuser","salt":"aMcmc5p1","md5":"321930e49c1e0797e27f97c5cffbb744","sha1":"f00f7f2f40a1b63c799d849dc7b1ac7a09124931","sha256":"8eaf321a3b77fe78de3124ce0d54969d14dc3e327e74801d826495baf6423b00"},"dob":{"date":"1955-09-06T14:26:49.215Z","age":66},"registered":{"date":"2014-02-27T05:56:03.362Z","age":7},"phone":"0461-2742163","cell":"0173-7395805","id":{"name":"","value":null},"picture":{"large":"https://randomuser.me/api/portraits/men/16.jpg","medium":"https://randomuser.me/api/portraits/med/men/16.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/16.jpg"},"nat":"DE"}`}},
		{"lt", `results[dob.age<30]`, []string{`{"gender":"male","name":{"title":"Mr","first":"Byron","last":"Cooper"},"location":{"street":{"number":7641,"name":"Lakeshore Rd"},"city":"St. Petersburg","state":"Colorado","country":"United States","postcode":85898,"coordinates":{"latitude":"-13.9778","longitude":"-137.0939"},"timezone":{"offset":"+8:00","description":"Beijing, Perth, Singapore, Hong Kong"}},"email":"byron.cooper@example.com","login":{"uuid":"40069ca0-fefd-4976-ab13-dff1c80bbf7e","username":"bigbird675","password":"nelson","salt":"h0ExMxtO","md5":"24dd7367f95c4377b2fef8e65cbd0c1b","sha1":"8c416504804eddc891da083eb3461733738fdc02","sha256":"96ce2b753601e405c00301242398c24551260208ac4799d78aacc2f8b87a7595"},"dob":{"date":"1992-02-22T10:15:55.954Z","age":29},"registered":{"date":"2007-03-13T17:08:10.687Z","age":14},"phone":"(230)-039-4521","cell":"(532)-586-9671","id":{"name":"SSN","value":"120-86-5372"},"picture":{"large":"https://randomuser.me/api/portraits/men/74.jpg","medium":"https://randomuser.me/api/portraits/med/men/74.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/74.jpg"},"nat":"US"}`}},
		{"lte", `results[dob.age<=31]`, []string{`{"gender":"male","name":{"title":"Mr","first":"Glen","last":"Howard"},"location":{"street":{"number":1180,"name":"George Street"},"city":"Swords","state":"Wexford","country":"Ireland","postcode":99620,"coordinates":{"latitude":"29.7085","longitude":"75.1163"},"timezone":{"offset":"+5:30","description":"Bombay, Calcutta, Madras, New Delhi"}},"email":"glen.howard@example.com","login":{"uuid":"a316c6e2-588b-46f7-bb84-b5259809db63","username":"greenbird880","password":"tttttt","salt":"gXFbTT9E","md5":"740e62a443776e531c32e8a27936cd32","sha1":"8cae36827b4626ee5d4bb33ec9a950b442c77be0","sha256":"299b67ef37fb0dfaee5069bf42cc4760cf4683b3f747ca5df974c2e22ef5940a"},"dob":{"date":"1990-02-13T12:02:30.019Z","age":31},"registered":{"date":"2009-06-27T06:37:56.602Z","age":12},"phone":"061-787-3251","cell":"081-991-2405","id":{"name":"PPS","value":"5290365T"},"picture":{"large":"https://randomuser.me/api/portraits/men/49.jpg","medium":"https://randomuser.me/api/portraits/med/men/49.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/49.jpg"},"nat":"IE"}`, `{"gender":"male","name":{"title":"Mr","first":"Byron","last":"Cooper"},"location":{"street":{"number":7641,"name":"Lakeshore Rd"},"city":"St. Petersburg","state":"Colorado","country":"United States","postcode":85898,"coordinates":{"latitude":"-13.9778","longitude":"-137.0939"},"timezone":{"offset":"+8:00","description":"Beijing, Perth, Singapore, Hong Kong"}},"email":"byron.cooper@example.com","login":{"uuid":"40069ca0-fefd-4976-ab13-dff1c80bbf7e","username":"bigbird675","password":"nelson","salt":"h0ExMxtO","md5":"24dd7367f95c4377b2fef8e65cbd0c1b","sha1":"8c416504804eddc891da083eb3461733738fdc02","sha256":"96ce2b753601e405c00301242398c24551260208ac4799d78aacc2f8b87a7595"},"dob":{"date":"1992-02-22T10:15:55.954Z","age":29},"registered":{"date":"2007-03-13T17:08:10.687Z","age":14},"phone":"(230)-039-4521","cell":"(532)-586-9671","id":{"name":"SSN","value":"120-86-5372"},"picture":{"large":"https://randomuser.me/api/portraits/men/74.jpg","medium":"https://randomuser.me/api/portraits/med/men/74.jpg","thumbnail":"https://randomuser.me/api/portraits/thumb/men/74.jpg"},"nat":"US"}`}},

		// slice
		{"dotSeperator", `results.name.[:1]`, []string{`{"title":"Mrs","first":"Melodie","last":"Gagné"}`}},
		{"noSeparator", `results.name[:1]`, []string{`{"title":"Mrs","first":"Melodie","last":"Gagné"}`}},
		{"last", `results.name[-1:]`, []string{`{"title":"Mr","first":"Koray","last":"Adal"}`}},

		// count
		{"count", `results.#`, []string{`10`}},
		{"countDown", `results.name.#`, []string{`10`}},
		{"countFilter", `results.name[title=Mrs].#`, []string{`3`}},
		{"countNotExists", `results.name[title=nothing].#`, []string{`0`}},

		// regex
		{"regexBasic", `results.name[first~^Y]`, []string{`{"title":"Miss","first":"Yesim","last":"Nijenkamp"}`}},
		{"regexContains", `results.name[last~N]`, []string{`{"title":"Miss","first":"Yesim","last":"Nijenkamp"}`}},
		{"regexSpecials01", `results.name[title~\w{4}]`, []string{`{"title":"Miss","first":"Yesim","last":"Nijenkamp"}`}},
	}
	for _, testcase := range testData {
		output, _ := ProcessExpression(testcase.input, testJson)
		var expectedTokens [][]byte
		for _, i := range testcase.expected {
			expectedTokens = append(expectedTokens, []byte(i))
		}
		if !reflect.DeepEqual(expectedTokens, output) {
			//t.Errorf("%s --> failed", testcase.name)
			t.Errorf("%s --> failed \n===\nexpected: %s\n===\nactual: %s\n", testcase.name, expectedTokens, output)
		}
	}

	// error cases
	errorData := []struct {
		name     string
		input    string
		expected common.ErrorCode
	}{
		{"closingParen", `results[name=test]]`, common.InvalidExpr},
		{"openParen", `results[[name=test]`, common.InvalidExpr},
		{"noMatchRegex", `results(name=test)`, common.InvalidExpr},
	}
	for _, testcase := range errorData {
		_, e := ProcessExpression(testcase.input, testJson)
		if e == nil || !strings.ContainsAny(e.Error(), testcase.expected.GetMsg()) {
			t.Errorf("%s --> failed \n===\nexpected: %s\n===\nactual: %s\n",
				testcase.name, testcase.expected.GetMsg(), e.Error())
		}
	}
}
