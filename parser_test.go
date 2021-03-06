package zipcode

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func ExampleParser_Parse() {
	tab := []string{
		`36207,"77703","7770301","ﾄｸｼﾏｹﾝ","ﾐﾏｼ","ｺﾔﾀﾞｲﾗ(ｲﾁﾊﾅ､ｲﾏﾏﾙ､ｵﾔﾏ､ｶｺﾞﾐ､ｶｼﾜﾗ､ｷﾅｶ､ｸﾜｶﾞﾗ､ｹﾔｷﾋﾗ､ｺﾋﾞｳﾗ､","徳島県","美馬市","木屋平（市初、今丸、尾山、カゴミ、樫原、木中、桑柄、ケヤキヒラ、小日浦、",1,0,0,0,0,0`,
		`36207,"77703","7770301","ﾄｸｼﾏｹﾝ","ﾐﾏｼ","ｼﾞｼﾞﾝﾀﾞｷ､ｽｹﾞｿﾞｳ､ﾂｴﾀﾞﾆ､ﾂﾂﾞﾛｳ､ﾊｼﾞｺﾉ､ﾋﾞﾔｶﾞｲﾁ､ﾌﾀﾄﾞ､ﾐﾂｷﾞ､ﾐﾂｸﾞ､","徳島県","美馬市","地神滝、菅蔵、杖谷、葛尾、ハジコノ、ビヤガイチ、二戸、三ツ木、貢、",1,0,0,0,0,0`,
		`36207,"77703","7770301","ﾄｸｼﾏｹﾝ","ﾐﾏｼ","ﾐﾅﾐﾊﾞﾘ､ﾑｺｳｶｼﾜﾗ)","徳島県","美馬市","南張、向樫原）",1,0,0,0,0,0`,
	}
	data := strings.Join(tab, "\r\n")
	fin := bytes.NewBufferString(data)

	var parser Parser
	c := parser.Parse(fin)
	var last *Entry
	for entry := range c {
		last = entry
	}
	if parser.Error != nil {
		log.Fatal(parser.Error)
	}
	fmt.Println(last.Pref.Text + last.Region.Text + last.Town.Text)
	// Output: 徳島県美馬市木屋平向樫原
}

func TestParse(t *testing.T) {
	actuals := []string{
		`01101,"060  ","0600000","ﾎｯｶｲﾄﾞｳ","ｻｯﾎﾟﾛｼﾁｭｳｵｳｸ","ｲｶﾆｹｲｻｲｶﾞﾅｲﾊﾞｱｲ","北海道","札幌市中央区","以下に掲載がない場合",0,0,0,0,0,0`,
		`13104,"160  ","1600023","ﾄｳｷｮｳﾄ","ｼﾝｼﾞｭｸｸ","ﾆｼｼﾝｼﾞｭｸ(ﾂｷﾞﾉﾋﾞﾙｦﾉｿﾞｸ)","東京都","新宿区","西新宿（次のビルを除く）",0,0,1,0,0,0`,
		`23105,"450  ","4506247","ｱｲﾁｹﾝ","ﾅｺﾞﾔｼﾅｶﾑﾗｸ","ﾒｲｴｷﾐｯﾄﾞﾗﾝﾄﾞｽｸｴｱ(ｺｳｿｳﾄｳ)(47ｶｲ)","愛知県","名古屋市中村区","名駅ミッドランドスクエア（高層棟）（４７階）",0,0,0,0,0,0`,
	}
	expects := []*Entry{
		&Entry{
			Code:   "01101",
			OldZip: "060  ",
			Zip:    "0600000",
			Pref:   Name{"北海道", "ﾎｯｶｲﾄﾞｳ"},
			Region: Name{"札幌市中央区", "ｻｯﾎﾟﾛｼﾁｭｳｵｳｸ"},
			Notice: "以下に掲載がない場合",
		},
		&Entry{
			Code:            "13104",
			OldZip:          "160  ",
			Zip:             "1600023",
			Pref:            Name{"東京都", "ﾄｳｷｮｳﾄ"},
			Region:          Name{"新宿区", "ｼﾝｼﾞｭｸｸ"},
			Town:            Name{"西新宿", "ﾆｼｼﾝｼﾞｭｸ"},
			IsBlockedScheme: true,
		},
		&Entry{
			Code:   "23105",
			OldZip: "450  ",
			Zip:    "4506247",
			Pref:   Name{"愛知県", "ｱｲﾁｹﾝ"},
			Region: Name{"名古屋市中村区", "ﾅｺﾞﾔｼﾅｶﾑﾗｸ"},
			Town:   Name{"名駅ミッドランドスクエア47階", "ﾒｲｴｷﾐｯﾄﾞﾗﾝﾄﾞｽｸｴｱ47ｶｲ"},
		},
	}
	parseTest(t, actuals, expects, "\n")
}

func TestParseSplitted(t *testing.T) {
	actuals := []string{
		`02206,"01855","0185501","ｱｵﾓﾘｹﾝ","ﾄﾜﾀﾞｼ","ｵｸｾ(ｱｵﾌﾞﾅ､ｺﾀﾀﾐｲｼ､ﾄﾜﾀﾞ､ﾄﾜﾀﾞｺﾊﾝｳﾀﾙﾍﾞ､ﾄﾜﾀﾞｺﾊﾝﾈﾉｸﾁ､","青森県","十和田市","奥瀬（青撫、小畳石、十和田、十和田湖畔宇樽部、十和田湖畔子ノ口、",1,1,0,0,0,0`,
		`02206,"01855","0185501","ｱｵﾓﾘｹﾝ","ﾄﾜﾀﾞｼ","ﾄﾜﾀﾞｺﾊﾝﾔｽﾐﾔ)","青森県","十和田市","十和田湖畔休屋）",1,1,0,0,0,0`,

		`26104,"604  ","6040983","ｷｮｳﾄﾌ","ｷｮｳﾄｼﾅｶｷﾞｮｳｸ","ｻｻﾔﾁｮｳ","京都府","京都市中京区","笹屋町（麩屋町通竹屋町下る、麩屋町通夷川上る、竹屋町通麩屋町西入、竹屋",0,0,0,0,0,0`,
		`26104,"604  ","6040983","ｷｮｳﾄﾌ","ｷｮｳﾄｼﾅｶｷﾞｮｳｸ","ｻｻﾔﾁｮｳ","京都府","京都市中京区","町通麩屋町東入、竹屋町通御幸町西入、夷川通麩屋町西入、夷川通麩屋町東入）",0,0,0,0,0,0`,
	}
	expects := []*Entry{
		&Entry{
			Code:          "02206",
			OldZip:        "01855",
			Zip:           "0185501",
			Pref:          Name{"青森県", "ｱｵﾓﾘｹﾝ"},
			Region:        Name{"十和田市", "ﾄﾜﾀﾞｼ"},
			Town:          Name{"奥瀬青撫", "ｵｸｾｱｵﾌﾞﾅ"},
			IsPartialTown: true,
			IsLargeTown:   true,
		},
		&Entry{
			Code:          "02206",
			OldZip:        "01855",
			Zip:           "0185501",
			Pref:          Name{"青森県", "ｱｵﾓﾘｹﾝ"},
			Region:        Name{"十和田市", "ﾄﾜﾀﾞｼ"},
			Town:          Name{"奥瀬小畳石", "ｵｸｾｺﾀﾀﾐｲｼ"},
			IsPartialTown: true,
			IsLargeTown:   true,
		},
		&Entry{
			Code:          "02206",
			OldZip:        "01855",
			Zip:           "0185501",
			Pref:          Name{"青森県", "ｱｵﾓﾘｹﾝ"},
			Region:        Name{"十和田市", "ﾄﾜﾀﾞｼ"},
			Town:          Name{"奥瀬十和田", "ｵｸｾﾄﾜﾀﾞ"},
			IsPartialTown: true,
			IsLargeTown:   true,
		},
		&Entry{
			Code:          "02206",
			OldZip:        "01855",
			Zip:           "0185501",
			Pref:          Name{"青森県", "ｱｵﾓﾘｹﾝ"},
			Region:        Name{"十和田市", "ﾄﾜﾀﾞｼ"},
			Town:          Name{"奥瀬十和田湖畔宇樽部", "ｵｸｾﾄﾜﾀﾞｺﾊﾝｳﾀﾙﾍﾞ"},
			IsPartialTown: true,
			IsLargeTown:   true,
		},
		&Entry{
			Code:          "02206",
			OldZip:        "01855",
			Zip:           "0185501",
			Pref:          Name{"青森県", "ｱｵﾓﾘｹﾝ"},
			Region:        Name{"十和田市", "ﾄﾜﾀﾞｼ"},
			Town:          Name{"奥瀬十和田湖畔子ノ口", "ｵｸｾﾄﾜﾀﾞｺﾊﾝﾈﾉｸﾁ"},
			IsPartialTown: true,
			IsLargeTown:   true,
		},
		&Entry{
			Code:          "02206",
			OldZip:        "01855",
			Zip:           "0185501",
			Pref:          Name{"青森県", "ｱｵﾓﾘｹﾝ"},
			Region:        Name{"十和田市", "ﾄﾜﾀﾞｼ"},
			Town:          Name{"奥瀬十和田湖畔休屋", "ｵｸｾﾄﾜﾀﾞｺﾊﾝﾔｽﾐﾔ"},
			IsPartialTown: true,
			IsLargeTown:   true,
		},

		&Entry{
			Code:   "26104",
			OldZip: "604  ",
			Zip:    "6040983",
			Pref:   Name{"京都府", "ｷｮｳﾄﾌ"},
			Region: Name{"京都市中京区", "ｷｮｳﾄｼﾅｶｷﾞｮｳｸ"},
			Town:   Name{"笹屋町麩屋町通竹屋町下る", "ｻｻﾔﾁｮｳ"},
		},
		&Entry{
			Code:   "26104",
			OldZip: "604  ",
			Zip:    "6040983",
			Pref:   Name{"京都府", "ｷｮｳﾄﾌ"},
			Region: Name{"京都市中京区", "ｷｮｳﾄｼﾅｶｷﾞｮｳｸ"},
			Town:   Name{"笹屋町麩屋町通夷川上る", "ｻｻﾔﾁｮｳ"},
		},
		&Entry{
			Code:   "26104",
			OldZip: "604  ",
			Zip:    "6040983",
			Pref:   Name{"京都府", "ｷｮｳﾄﾌ"},
			Region: Name{"京都市中京区", "ｷｮｳﾄｼﾅｶｷﾞｮｳｸ"},
			Town:   Name{"笹屋町竹屋町通麩屋町西入", "ｻｻﾔﾁｮｳ"},
		},
		&Entry{
			Code:   "26104",
			OldZip: "604  ",
			Zip:    "6040983",
			Pref:   Name{"京都府", "ｷｮｳﾄﾌ"},
			Region: Name{"京都市中京区", "ｷｮｳﾄｼﾅｶｷﾞｮｳｸ"},
			Town:   Name{"笹屋町竹屋町通麩屋町東入", "ｻｻﾔﾁｮｳ"},
		},
		&Entry{
			Code:   "26104",
			OldZip: "604  ",
			Zip:    "6040983",
			Pref:   Name{"京都府", "ｷｮｳﾄﾌ"},
			Region: Name{"京都市中京区", "ｷｮｳﾄｼﾅｶｷﾞｮｳｸ"},
			Town:   Name{"笹屋町竹屋町通御幸町西入", "ｻｻﾔﾁｮｳ"},
		},
		&Entry{
			Code:   "26104",
			OldZip: "604  ",
			Zip:    "6040983",
			Pref:   Name{"京都府", "ｷｮｳﾄﾌ"},
			Region: Name{"京都市中京区", "ｷｮｳﾄｼﾅｶｷﾞｮｳｸ"},
			Town:   Name{"笹屋町夷川通麩屋町西入", "ｻｻﾔﾁｮｳ"},
		},
		&Entry{
			Code:   "26104",
			OldZip: "604  ",
			Zip:    "6040983",
			Pref:   Name{"京都府", "ｷｮｳﾄﾌ"},
			Region: Name{"京都市中京区", "ｷｮｳﾄｼﾅｶｷﾞｮｳｸ"},
			Town:   Name{"笹屋町夷川通麩屋町東入", "ｻｻﾔﾁｮｳ"},
		},
	}
	parseTest(t, actuals, expects, "\n")
}

func TestPraseExpr(t *testing.T) {
	actuals := []string{
		`01101,"064  ","0640930","ﾎｯｶｲﾄﾞｳ","ｻｯﾎﾟﾛｼﾁｭｳｵｳｸ","ﾐﾅﾐ30ｼﾞｮｳﾆｼ(9-11ﾁｮｳﾒ)","北海道","札幌市中央区","南三十条西（９〜１１丁目）",0,0,1,0,0,0`,

		// TODO: xxx-xxが含まれるパターンを実装
		//`01303,"06137","0613774","ﾎｯｶｲﾄﾞｳ","ｲｼｶﾘｸﾞﾝﾄｳﾍﾞﾂﾁｮｳ","ｶﾜｼﾓ(782-13､5363-7-8､5382-3､5405","北海道","石狩郡当別町","川下（７８２−１３、５３６３−７〜８、５３８２−３、５４０５",1,0,0,0,0,0`,
		//`01303,"06137","0613774","ﾎｯｶｲﾄﾞｳ","ｲｼｶﾘｸﾞﾝﾄｳﾍﾞﾂﾁｮｳ","-4､5407-5､5445-5446-4ﾊﾞﾝﾁ)","北海道","石狩郡当別町","−４、５４０７−５、５４４５〜５４４６−４番地）",1,0,0,0,0,0`,

		// TODO: 単純な地割
		// `03202,"02824","0282402","ｲﾜﾃｹﾝ","ﾐﾔｺｼ","ｶﾜｲ(ﾀﾞｲ9ﾁﾜﾘ-ﾀﾞｲ11ﾁﾜﾘ)","岩手県","宮古市","川井（第９地割〜第１１地割）",1,1,0,0,0,0`

		// TODO: 複数の地割
		// `03202,"02825","0282504","ｲﾜﾃｹﾝ","ﾐﾔｺｼ","ﾊｺｲｼ(ﾀﾞｲ2ﾁﾜﾘ<70-136>-ﾀﾞｲ4ﾁﾜﾘ<3-11>)","岩手県","宮古>市","箱石（第２地割「７０〜１３６」〜第４地割「３〜１１」）",1,1,0,0,0,0`,

		// TODO: 地番と除くの組み合わせ実装
		//`03302,"02851","0285102","ｲﾜﾃｹﾝ","ｲﾜﾃｸﾞﾝｸｽﾞﾏｷﾏﾁ","ｸｽﾞﾏｷ(ﾀﾞｲ40ﾁﾜﾘ<57ﾊﾞﾝﾁ125､176ｦﾉｿﾞｸ>-ﾀﾞｲ45","岩手県","岩手郡葛巻町","葛巻（第４０地割「５７番地１２５、１７６を除く」〜第４５",1,1,0,0,0,0`,
		//`03302,"02851","0285102","ｲﾜﾃｹﾝ","ｲﾜﾃｸﾞﾝｸｽﾞﾏｷﾏﾁ","ﾁﾜﾘ)","岩手県","岩手郡葛巻町","地割）",1,1,0,0,0,0`,
	}
	expects := []*Entry{
		&Entry{
			Code:            "01101",
			OldZip:          "064  ",
			Zip:             "0640930",
			Pref:            Name{"北海道", "ﾎｯｶｲﾄﾞｳ"},
			Region:          Name{"札幌市中央区", "ｻｯﾎﾟﾛｼﾁｭｳｵｳｸ"},
			Town:            Name{"南三十条西9丁目", "ﾐﾅﾐ30ｼﾞｮｳﾆｼ9ﾁｮｳﾒ"},
			IsBlockedScheme: true,
		},
		&Entry{
			Code:            "01101",
			OldZip:          "064  ",
			Zip:             "0640930",
			Pref:            Name{"北海道", "ﾎｯｶｲﾄﾞｳ"},
			Region:          Name{"札幌市中央区", "ｻｯﾎﾟﾛｼﾁｭｳｵｳｸ"},
			Town:            Name{"南三十条西10丁目", "ﾐﾅﾐ30ｼﾞｮｳﾆｼ10ﾁｮｳﾒ"},
			IsBlockedScheme: true,
		},
		&Entry{
			Code:            "01101",
			OldZip:          "064  ",
			Zip:             "0640930",
			Pref:            Name{"北海道", "ﾎｯｶｲﾄﾞｳ"},
			Region:          Name{"札幌市中央区", "ｻｯﾎﾟﾛｼﾁｭｳｵｳｸ"},
			Town:            Name{"南三十条西11丁目", "ﾐﾅﾐ30ｼﾞｮｳﾆｼ11ﾁｮｳﾒ"},
			IsBlockedScheme: true,
		},
	}
	parseTest(t, actuals, expects, "\n")

}

func TestParseOthers(t *testing.T) {
	actuals := []string{
		`02206,"03403","0340301","ｱｵﾓﾘｹﾝ","ﾄﾜﾀﾞｼ","ｵｸｾ(ｿﾉﾀ)","青森県","十和田市","奥瀬（その他）",1,1,0,0,0,0`,
	}
	expects := []*Entry{
		&Entry{
			Code:          "02206",
			OldZip:        "03403",
			Zip:           "0340301",
			Pref:          Name{"青森県", "ｱｵﾓﾘｹﾝ"},
			Region:        Name{"十和田市", "ﾄﾜﾀﾞｼ"},
			Town:          Name{"奥瀬", "ｵｸｾ"},
			IsPartialTown: true,
			IsLargeTown:   true,
		},
	}
	parseTest(t, actuals, expects, "\n")
}

func TestPraseRegionWithNumber(t *testing.T) {
	actuals := []string{
		`38204,"796  ","7960088","ｴﾋﾒｹﾝ","ﾔﾜﾀﾊﾏｼ","ﾔﾜﾀﾊﾏｼﾉﾂｷﾞﾆﾊﾞﾝﾁｶﾞｸﾙﾊﾞｱｲ","愛媛県","八幡浜市","八幡浜市の次に番地がくる場合",0,0,0,0,0,0`,
		`39386,"78121","7812110","ｺｳﾁｹﾝ","ｱｶﾞﾜｸﾞﾝｲﾉﾁｮｳ","ｲﾉﾁｮｳﾉﾂｷﾞﾆﾊﾞﾝﾁｶﾞｸﾙﾊﾞｱｲ","高知県","吾川郡いの町","いの町の次に番地がくる場合",0,0,0,0,0,0`,
		`42212,"85724","8572427","ﾅｶﾞｻｷｹﾝ","ｻｲｶｲｼ","ｵｵｼﾏﾁｮｳﾉﾂｷﾞﾆﾊﾞﾝﾁｶﾞｸﾙﾊﾞｱｲ","長崎県","西海市","大島町の次に番地がくる場合",0,0,0,0,0,0`,
	}
	expects := []*Entry{
		&Entry{
			Code:   "38204",
			OldZip: "796  ",
			Zip:    "7960088",
			Pref:   Name{"愛媛県", "ｴﾋﾒｹﾝ"},
			Region: Name{"八幡浜市", "ﾔﾜﾀﾊﾏｼ"},
			Notice: "八幡浜市の次に番地がくる場合",
		},
		&Entry{
			Code:   "39386",
			OldZip: "78121",
			Zip:    "7812110",
			Pref:   Name{"高知県", "ｺｳﾁｹﾝ"},
			Region: Name{"吾川郡いの町", "ｱｶﾞﾜｸﾞﾝｲﾉﾁｮｳ"},
			Notice: "いの町の次に番地がくる場合",
		},
		&Entry{
			Code:   "42212",
			OldZip: "85724",
			Zip:    "8572427",
			Pref:   Name{"長崎県", "ﾅｶﾞｻｷｹﾝ"},
			Region: Name{"西海市", "ｻｲｶｲｼ"},
			Town:   Name{"大島町", "ｵｵｼﾏﾁｮｳ"},
			Notice: "大島町の次に番地がくる場合",
		},
	}
	parseTest(t, actuals, expects, "\n")
}

func TestPraseCircle(t *testing.T) {
	actuals := []string{
		`13362,"10003","1000301","ﾄｳｷｮｳﾄ","ﾄｼﾏﾑﾗ","ﾄｼﾏﾑﾗｲﾁｴﾝ","東京都","利島村","利島村一円",0,0,0,0,0,0`,
		`25443,"52203","5220317","ｼｶﾞｹﾝ","ｲﾇｶﾐｸﾞﾝﾀｶﾞﾁｮｳ","ｲﾁｴﾝ","滋賀県","犬上郡多賀町","一円",0,0,0,0,0,0`,
	}
	expects := []*Entry{
		&Entry{
			Code:   "13362",
			OldZip: "10003",
			Zip:    "1000301",
			Pref:   Name{"東京都", "ﾄｳｷｮｳﾄ"},
			Region: Name{"利島村", "ﾄｼﾏﾑﾗ"},
			Notice: "利島村一円",
		},
		&Entry{
			Code:   "25443",
			OldZip: "52203",
			Zip:    "5220317",
			Pref:   Name{"滋賀県", "ｼｶﾞｹﾝ"},
			Region: Name{"犬上郡多賀町", "ｲﾇｶﾐｸﾞﾝﾀｶﾞﾁｮｳ"},
			Town:   Name{"一円", "ｲﾁｴﾝ"},
		},
	}
	parseTest(t, actuals, expects, "\n")
}

func parseTest(t *testing.T, actuals []string, expects []*Entry, newline string) {
	s := strings.Join(actuals, newline)
	fin := bytes.NewBufferString(s)

	var parser Parser
	c := parser.Parse(fin)
	for _, expect := range expects {
		entry := <-c
		if entry == nil {
			t.Errorf("Parse() = nil; Expect %v", expect)
			continue
		}
		if entry.Code != expect.Code {
			t.Errorf("Parse(): Code = %q; Expect %q", entry.Code, expect.Code)
		}
		if entry.OldZip != expect.OldZip {
			t.Errorf("Parse(): OldZip = %q; Expect %q", entry.OldZip, expect.OldZip)
		}
		if entry.Zip != expect.Zip {
			t.Errorf("Parse(): Zip = %q; Expect %q", entry.Zip, expect.Zip)
		}
		if !entry.Pref.Equal(expect.Pref) {
			t.Errorf("Parse(): Pref = %q; Expect %q", entry.Pref, expect.Pref)
		}
		if !entry.Region.Equal(expect.Region) {
			t.Errorf("Parse(): Region = %q; Expect %q", entry.Region, expect.Region)
		}
		if !entry.Town.Equal(expect.Town) {
			t.Errorf("Parse(): Town = %q; Expect %q", entry.Town, expect.Town)
		}
		if entry.IsPartialTown != expect.IsPartialTown {
			t.Errorf("Parse(): IsPartialTown = %t; Expect %t", entry.IsPartialTown, expect.IsPartialTown)
		}
		if entry.IsLargeTown != expect.IsLargeTown {
			t.Errorf("Parse(): IsLargeTown = %t; Expect %t", entry.IsLargeTown, expect.IsLargeTown)
		}
		if entry.IsBlockedScheme != expect.IsBlockedScheme {
			t.Errorf("Parse(): IsBlockedScheme = %t; Expect %t", entry.IsBlockedScheme, expect.IsBlockedScheme)
		}
		if entry.IsOverlappedZip != expect.IsOverlappedZip {
			t.Errorf("Parse(): IsOverlappedZip = %t; Expect %t", entry.IsOverlappedZip, expect.IsOverlappedZip)
		}
		if entry.Status != expect.Status {
			t.Errorf("Parse(): Status = %q; Expect %q", entry.Status, expect.Status)
		}
		if entry.Reason != expect.Reason {
			t.Errorf("Parse(): Reason = %q; Expect %q", entry.Reason, expect.Reason)
		}
		if entry.Notice != expect.Notice {
			t.Errorf("Parse(): Notice = %q; Expect %q", entry.Notice, expect.Notice)
		}
	}
	if entry, ok := <-c; ok {
		t.Errorf("Parse() = %v; Expect end", *entry)
	}
	if parser.Error != nil {
		t.Fatalf("Parse() = %v; Expect not error", parser.Error)
	}
}
