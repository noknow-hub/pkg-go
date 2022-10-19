//////////////////////////////////////////////////////////////////////
// add_client.go
//////////////////////////////////////////////////////////////////////
package country

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    myQuery "github.com/noknow-hub/pkg-go/db/mysql/query"
)

type AddClient struct {
    BaseClient *myQuery.InsertClient
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDb(tableName string, db *sql.DB) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithDb(tableName, db),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with db object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithDbContext(tableName string, db *sql.DB, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithDbContext(tableName, db, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTx(tableName string, tx *sql.Tx) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithTx(tableName, tx),
    }
}


//////////////////////////////////////////////////////////////////////
// New AddClient with tx object and context.
//////////////////////////////////////////////////////////////////////
func NewAddClientWithTxContext(tableName string, tx *sql.Tx, ctx context.Context) *AddClient {
    return &AddClient{
        BaseClient: myQuery.NewInsertClientWithTxContext(tableName, tx, ctx),
    }
}


//////////////////////////////////////////////////////////////////////
// Run.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) Run() (*myQuery.InsertResult, error) {
    return c.BaseClient.Run()
}


//////////////////////////////////////////////////////////////////////
// Run for initialization.
//////////////////////////////////////////////////////////////////////
func (c *AddClient) RunForInitialization() (*myQuery.InsertResult, error) {
    c.BaseClient.SetIgnore().
            SetColNames([]string{COL_COUNTRY_CODE, COL_AR, COL_DE, COL_EN, COL_ES, COL_FR, COL_JA, COL_PT, COL_RU, COL_ZH_CN, COL_ZH_TW, COL_CONTINENT, COL_STATUS}).
            AppendValues([]interface{}{"ad","أندورا","Andorra","Andorra","Andorra","Andorre","アンドラ","Andorra","Андорра","安道尔","安道尔",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ae","الإمارات العربية المتحدة","Vereinigte Arabische Emirate","United Arab Emirates","Emiratos Árabes Unidos","Émirats arabes unis","アラブ首長国連邦","Emirados Árabes Unidos","ОАЭ","阿联酋","阿联酋",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"af","أفغانستان","Afghanistan","Afghanistan","Afganistán","Afghanistan","アフガニスタン","Afeganistão","Афганистан","阿富汗","阿富汗",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ag","أنتيغوا وباربودا","Antigua und Barbuda","Antigua and Barbuda","Antigua y Barbuda","Antigua-et-Barbuda","アンティグア・バーブーダ","Antígua e Barbuda","Антигуа и Барбуда","安地卡及巴布達","安地卡及巴布達",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ai","أنغويلا","Anguilla","Anguilla","Anguila","Anguilla","アンギラ","Anguilla","Ангилья","安圭拉","安圭拉",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"al","ألبانيا","Albanien","Albania","Albania","Albanie","アルバニア","Albânia","Албания","阿尔巴尼亚","阿尔巴尼亚",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"am","أرمينيا","Armenien","Armenia","Armenia","Arménie","アルメニア","Armênia","Армения","亞美尼亞","亞美尼亞",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ao","أنغولا","Angola","Angola","Angola","Angola","アンゴラ","Angola","Ангола","安哥拉","安哥拉",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"aq","القارة القطبية الجنوبية","Antarktika","Antarctica","Antártida","Antarctique","南極","Antártica","Антарктида","南极洲","南极洲",7,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ar","الأرجنتين","Argentinien","Argentina","Argentina","Argentine","アルゼンチン","Argentina","Аргентина","阿根廷","阿根廷",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"as","ساموا الأمريكية","Amerikanisch-Samoa","American Samoa","Samoa Americana","Samoa américaines","アメリカ領サモア","Samoa Americana","Американское Самоа","美属萨摩亚","美属萨摩亚",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"at","النمسا","Österreich","Austria","Austria","Autriche","オーストリア","Áustria","Австрия","奥地利","奥地利",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"au","أستراليا","Australien","Australia","Australia","Australie","オーストラリア","Austrália","Австралия","澳大利亚","澳大利亚",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"aw","أروبا","Aruba","Aruba","Aruba","Aruba","アルバ","Aruba","Аруба","阿鲁巴","阿鲁巴",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ax","جزر أولاند","Åland","Åland Islands","Åland","les Åland","オーランド諸島","Ilhas Aland","Аландские острова","奥兰","奥兰",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"az","أذربيجان","Aserbaidschan","Azerbaijan","Azerbaiyán","Azerbaïdjan","アゼルバイジャン","Azerbaijão","Азербайджан","阿塞拜疆","阿塞拜疆",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ba","البوسنة والهرسك","Bosnien und Herzegowina","Bosnia and Herzegovina","Bosnia y Herzegovina","Bosnie-Herzégovine","ボスニア・ヘルツェゴビナ","Bósnia e Herzegovina","Босния и Герцеговина","波斯尼亚和黑塞哥维那","波斯尼亚和黑塞哥维那",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bb","باربادوس","Barbados","Barbados","Barbados","Barbade","バルバドス","Barbados","Барбадос","巴巴多斯","巴巴多斯",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bd","بنغلاديش","Bangladesch","Bangladesh","Bangladés","Bangladesh","バングラデシュ","Bangladesh","Бангладеш","孟加拉国","孟加拉国",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"be","بلجيكا","Belgien","Belgium","Bélgica","Belgique","ベルギー","Bélgica","Бельгия","比利時","比利時",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bf","بوركينا فاسو","Burkina Faso","Burkina Faso","Burkina Faso","Burkina Faso","ブルキナファソ","Burkina Faso","Буркина-Фасо","布吉納法索","布吉納法索",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bg","بلغاريا","Bulgarien","Bulgaria","Bulgaria","Bulgarie","ブルガリア","Bulgária","Болгария","保加利亚","保加利亚",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bh","البحرين","Bahrain","Bahrain","Baréin","Bahreïn","バーレーン","Barém","Бахрейн","巴林","巴林",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bi","بوروندي","Burundi","Burundi","Burundi","Burundi","ブルンジ","Burundi","Бурунди","布隆迪","布隆迪",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bj","بنين","Benin","Benin","Benín","Bénin","ベナン","Benin","Бенин","贝宁","贝宁",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bl","سان بارتيلمي","Saint-Barthélemy","Saint Barthélemy","San Bartolomé","Saint-Barthélemy","サン・バルテルミー","São Bartolomeu","Сен-Бартелеми","圣巴泰勒米","圣巴泰勒米",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bm","برمودا","Bermuda","Bermuda","Bermudas","Bermudes","バミューダ","Bermudas","Бермуды","百慕大","百慕大",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bn","بروناي","Brunei Darussalam","Brunei Darussalam","Brunéi","Brunei","ブルネイ・ダルサラーム","Brunei Darussalam","Бруней","文莱","文莱",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bo","بوليفيا","Bolivien","Bolivia","Bolivia","Bolivie","ボリビア多民族国","Bolívia","Боливия","玻利维亚","玻利维亚",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bq","الجزر الكاريبية الهولندية","Bonaire","Bonaire","Bonaire","Pays-Bas caribéens","ボネール","Bonaire","Бонэйр, Синт-Эстатиус и Саба","荷兰加勒比区","荷兰加勒比区",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"br","البرازيل","Brasilien","Brazil","Brasil","Brésil","ブラジル","Brasil","Бразилия","巴西","巴西",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bs","باهاماس","Bahamas","Bahamas","Bahamas","Bahamas","バハマ","Bahamas","Багамские Острова","巴哈马","巴哈马",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bt","بوتان","Bhutan","Bhutan","Bután","Bhoutan","ブータン","Butão","Бутан","不丹","不丹",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bv","جزيرة بوفيه","Bouvetinsel","Bouvet Island","Isla Bouvet","Île Bouvet","ブーベ島","Ilha Bouvet","Остров Буве","布韦岛","布韦岛",7,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bw","بوتسوانا","Botswana","Botswana","Botsuana","Botswana","ボツワナ","Botsuana","Ботсвана","博茨瓦纳","博茨瓦纳",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"by","روسيا البيضاء","Belarus","Belarus","Bielorrusia","Biélorussie","ベラルーシ","Bielorrússia","Белоруссия","白俄羅斯","白俄羅斯",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"bz","بليز","Belize","Belize","Belice","Belize","ベリーズ","Belize","Белиз","伯利兹","伯利兹",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ca","كندا","Kanada","Canada","Canadá","Canada","カナダ","Canadá","Канада","加拿大","加拿大",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cc","جزر كوكوس","Kokosinseln","Cocos (Keeling) Islands","Islas Cocos","Îles Cocos","ココス (キーリング) 諸島","Ilhas Cocos","Кокосовые острова","科科斯 (基林) 群島","科科斯 (基林) 群島",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cd","جمهورية الكونغو الديمقراطية","Kongo, Demokratische Republik","Congo, Democratic Republic of the","República Democrática del Congo","République démocratique du Congo","コンゴ民主共和国","Congo, República Democrática do","ДР Конго","刚果 (金)","刚果 (金)",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cf","جمهورية أفريقيا الوسطى","Zentralafrikanische Republik","Central African Republic","República Centroafricana","République centrafricaine","中央アフリカ共和国","República Centro-Africana","ЦАР","中非","中非",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cg","جمهورية الكونغو","Kongo, Republik","Congo","República del Congo","République du Congo","コンゴ共和国","Congo","Республика Конго","刚果 (布)","",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ch","سويسرا","Schweiz","Switzerland","Suiza","Suisse","スイス","Suíço","Швейцария","瑞士","瑞士",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ci","ساحل العاج","Côte d''Ivoire","Côte d''Ivoire","Costa de Marfil","Côte d''Ivoire","コートジボワール","Costa do Marfim","Кот-д’Ивуар","科特迪瓦","科特迪瓦",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ck","جزر كوك","Cookinseln","Cook Islands","Islas Cook","Îles Cook","クック諸島","Ilhas Cook","Острова Кука","庫克群島","庫克群島",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cl","تشيلي","Chile","Chile","Chile","Chili","チリ","Chile","Чили","智利","智利",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cm","الكاميرون","Kamerun","Cameroon","Camerún","Cameroun","カメルーン","Camarões","Камерун","喀麦隆","喀麦隆",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cn","الصين","China","China","China","Chine","中華人民共和国","China","Китай","中国","中国",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"co","كولومبيا","Kolumbien","Colombia","Colombia","Colombie","コロンビア","Colômbia","Колумбия","哥伦比亚","哥伦比亚",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cr","كوستاريكا","Costa Rica","Costa Rica","Costa Rica","Costa Rica","コスタリカ","Costa Rica","Коста-Рика","哥斯达黎加","哥斯达黎加",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cu","كوبا","Kuba","Cuba","Cuba","Cuba","キューバ","Cuba","Куба","古巴","古巴",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cv","الرأس الأخضر","Kap Verde","Cabo Verde","Cabo Verde","Cap-Vert","カーボベルデ","Cabo Verde","Кабо-Верде","佛得角","佛得角",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cw","كوراساو","Curaçao","Curaçao","Curazao","Curaçao","キュラソー","Curaçao","Кюрасао","库拉索","库拉索",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cx","جزيرة عيد الميلاد","Weihnachtsinsel","Christmas Island","Isla de Navidad","Île Christmas","クリスマス島","Ilha do Natal","Остров Рождества","圣诞岛","圣诞岛",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cy","قبرص","Zypern","Cyprus","Chipre","Chypre","キプロス","Chipre","Кипр","賽普勒斯","賽普勒斯",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"cz","جمهورية التشيك","Tschechien","Czech","República Checa","Tchéquie","チェコ","Tcheca","Чехия","捷克","捷克",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"de","ألمانيا","Deutschland","Germany","Alemania","Allemagne","ドイツ","Alemanha","Германия","德國","德國",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"dj","جيبوتي","Dschibuti","Djibouti","Yibuti","Djibouti","ジブチ","Djibuti","Джибути","吉布提","吉布提",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"dk","الدنمارك","Dänemark","Denmark","Dinamarca","Danemark","デンマーク","Dinamarca","Дания","丹麥","丹麥",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"dm","دومينيكا","Dominica","Dominica","Dominica","Dominique","ドミニカ国","Dominica","Доминика","多米尼克","多米尼克",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"do","جمهورية الدومينيكان","Dominikanische Republik","Dominican Republic","República Dominicana","République dominicaine","ドミニカ共和国","República Dominicana","Доминиканская Республика","多米尼加","多米尼加",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"dz","الجزائر","Algerien","Algeria","Argelia","Algérie","アルジェリア","Argélia","Алжир","阿尔及利亚","阿尔及利亚",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ec","الإكوادور","Ecuador","Ecuador","Ecuador","Équateur","エクアドル","Equador","Эквадор","厄瓜多尔","厄瓜多尔",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ee","إستونيا","Estland","Estonia","Estonia","Estonie","エストニア","Estônia","Эстония","爱沙尼亚","爱沙尼亚",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"eg","مصر","Ägypten","Egypt","Egipto","Égypte","エジプト","Egito","Египет","埃及","埃及",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"eh","الصحراء الغربية","Westsahara","Western Sahara","República Árabe Saharaui Democrática","République arabe sahraouie démocratique","西サハラ","Saara Ocidental","САДР","西撒哈拉","西撒哈拉",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"er","إريتريا","Eritrea","Eritrea","Eritrea","Érythrée","エリトリア","Eritreia","Эритрея","厄立特里亚","厄立特里亚",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"es","إسبانيا","Spanien","Spain","España","Espagne","スペイン","Espanha","Испания","西班牙","西班牙",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"et","إثيوبيا","Äthiopien","Ethiopia","Etiopía","Éthiopie","エチオピア","Etiópia","Эфиопия","衣索比亞","衣索比亞",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"fi","فنلندا","Finnland","Finland","Finlandia","Finlande","フィンランド","Finlândia","Финляндия","芬兰","芬兰",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"fj","فيجي","Fidschi","Fiji","Fiyi","Fidji","フィジー","Fiji","Фиджи","斐济","斐济",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"fk","جزر فوكلاند","Falklandinseln","Falkland Islands (Malvinas)","Islas Malvinas","Malouines","フォークランド (マルビナス) 諸島","Ilhas Falkland (Malvinas)","Фолклендские острова","福克蘭群島","福克蘭群島",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"fm","ولايات ميكرونيسيا المتحدة","Mikronesien","Micronesia (Federated States of)","Micronesia","États fédérés de Micronésie","ミクロネシア連邦","Micronésia (Estados Federados da)","Микронезия","密克羅尼西亞聯邦","密克羅尼西亞聯邦",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"fo","جزر فارو","Färöer","Faroe Islands","Islas Feroe","Îles Féroé","フェロー諸島","ilhas Faroe","Фареры","法罗群岛","法罗群岛",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"fr","فرنسا","Frankreich","France","Francia","France","フランス","França","Франция","法国","法国",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ga","الغابون","Gabun","Gabon","Gabón","Gabon","ガボン","Gabão","Габон","加彭","加彭",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gb","المملكة المتحدة","Vereinigtes Königreich Großbritannien und Nordirland","United Kingdom","Reino Unido","Royaume-Uni","イギリス","Reino Unido","Великобритания","英國","英國",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gd","غرينادا","Grenada","Grenada","Granada","Grenade","グレナダ","Granada","Гренада","格瑞那達","格瑞那達",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ge","جورجيا","Georgien","Georgia","Georgia","Géorgie","ジョージア","Geórgia","Грузия","格鲁吉亚","格鲁吉亚",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gf","غويانا الفرنسية","Französisch-Guayana","French Guiana","Guayana Francesa","Guyane","フランス領ギアナ","Guiana Francesa","Гвиана","法属圭亚那","法属圭亚那",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gg","غيرنزي","Guernsey","Guernsey","Guernsey","Guernesey","ガーンジー","Guernsey","Гернси","根西","根西",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gh","غانا","Ghana","Ghana","Ghana","Ghana","ガーナ","Gana","Гана","加纳","加纳",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gi","جبل طارق","Gibraltar","Gibraltar","Gibraltar","Gibraltar","ジブラルタル","Gibraltar","Гибралтар","直布罗陀","直布罗陀",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gl","جرينلاند","Grönland","Greenland","Groenlandia","Groenland","グリーンランド","Gronelândia","Гренландия","格陵兰","格陵兰",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gm","غامبيا","Gambia","Gambia","Gambia","Gambie","ガンビア","Gâmbia","Гамбия","冈比亚","冈比亚",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gn","غينيا","Guinea","Guinea","Guinea","Guinée","ギニア","Guiné","Гвинея","几内亚","几内亚",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gp","غوادلوب","Guadeloupe","Guadeloupe","Guadalupe","Guadeloupe","グアドループ","Guadalupe","Гваделупа","瓜德罗普","瓜德罗普",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gq","غينيا الاستوائية","Äquatorialguinea","Equatorial Guinea","Guinea Ecuatorial","Guinée équatoriale","赤道ギニア","Guiné Equatorial","Экваториальная Гвинея","赤道几内亚","赤道几内亚",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gr","اليونان","Griechenland","Greece","Grecia","Grèce","ギリシャ","Grécia","Греция","希臘","希臘",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gs","جورجيا الجنوبية وجزر ساندويتش الجنوبية","Südgeorgien und die Südlichen Sandwichinseln","South Georgia and the South Sandwich Islands","Islas Georgias del Sur y Sandwich del Sur","Géorgie du Sud-et-les îles Sandwich du Sud","サウスジョージア・サウスサンドウィッチ諸島","Ilhas Geórgia do Sul e Sandwich do Sul","Южная Георгия и Южные Сандвичевы Острова","南乔治亚和南桑威奇群岛","南乔治亚和南桑威奇群岛",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gt","غواتيمالا","Guatemala","Guatemala","Guatemala","Guatemala","グアテマラ","Guatemala","Гватемала","危地马拉","危地马拉",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gu","غوام","Guam","Guam","Guam","Guam","グアム","Guam","Гуам","關島","關島",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gw","غينيا بيساو","Guinea-Bissau","Guinea-Bissau","Guinea-Bisáu","Guinée-Bissau","ギニアビサウ","Guiné-Bissau","Гвинея-Бисау","几内亚比绍","几内亚比绍",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"gy","غيانا","Guyana","Guyana","Guyana","Guyana","ガイアナ","Guiana","Гайана","圭亚那","圭亚那",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"hk","هونغ كونغ","Hongkong","Hong Kong","Hong Kong","Hong Kong","香港","Hong Kong","Гонконг","香港","香港",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"hm","جزيرة هيرد وجزر ماكدونالد","Heard und McDonaldinseln","Heard Island and McDonald Islands","Islas Heard y McDonald","Îles Heard-et-MacDonald","ハード島とマクドナルド諸島","Ilha Heard e Ilhas McDonald","Херд и Макдональд","赫德岛和麦克唐纳群岛","赫德岛和麦克唐纳群岛",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"hn","هندوراس","Honduras","Honduras","Honduras","Honduras","ホンジュラス","Honduras","Гондурас","洪都拉斯","洪都拉斯",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"hr","كرواتيا","Kroatien","Croatia","Croacia","Croatie","クロアチア","Croácia","Хорватия","克罗地亚","克罗地亚",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ht","هايتي","Haiti","Haiti","Haití","Haïti","ハイチ","Haiti","Гаити","海地","海地",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"hu","المجر","Ungarn","Hungary","Hungría","Hongrie","ハンガリー","Hungria","Венгрия","匈牙利","匈牙利",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"id","إندونيسيا","Indonesien","Indonesia","Indonesia","Indonésie","インドネシア","Indonésia","Индонезия","印尼","印尼",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ie","أيرلندا","Irland","Ireland","Irlanda","Irlande","アイルランド","Irlanda","Ирландия","爱尔兰","爱尔兰",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"il","إسرائيل","Israel","Israel","Israel","Israël","イスラエル","Israel","Израиль","以色列","以色列",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"im","جزيرة مان","Insel Man","Isle of Man","Isla de Man","Île de Man","マン島","Ilha de Man","Остров Мэн","马恩岛","马恩岛",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"in","الهند","Indien","India","India","Inde","インド","Índia","Индия","印度","印度",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"io","إقليم المحيط الهندي البريطاني","Britisches Territorium im Indischen Ozean","British Indian Ocean Territory","Territorio Británico del Océano Índico","Territoire britannique de l''océan Indien","イギリス領インド洋地域","Território Britânico do Oceano Índico","Британская территория в Индийском океане","英屬印度洋領地","英屬印度洋領地",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"iq","العراق","Irak","Iraq","Irak","Irak","イラク","Iraque","Ирак","伊拉克","伊拉克",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ir","إيران","Iran, Islamische Republik","Iran (Islamic Republic of)","Irán","Iran","イラン・イスラム共和国","Irã (Republic Islâmica do Irã)","Иран","伊朗","伊朗",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"is","آيسلندا","Island","Iceland","Islandia","Islande","アイスランド","Islândia","Исландия","冰島","冰島",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"it","إيطاليا","Italien","Italy","Italia","Italie","イタリア","Itália","Италия","義大利","義大利",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"je","جيرزي","Jersey","Jersey","Jersey","Jersey","ジャージー","Jersey","Джерси","澤西","澤西",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"jm","جامايكا","Jamaika","Jamaica","Jamaica","Jamaïque","ジャマイカ","Jamaica","Ямайка","牙买加","牙买加",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"jo","الأردن","Jordanien","Jordan","Jordania","Jordanie","ヨルダン","Jordânia","Иордания","约旦","约旦",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"jp","اليابان","Japan","Japan","Japón","Japon","日本","Japão","Япония","日本","日本",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ke","كينيا","Kenia","Kenya","Kenia","Kenya","ケニア","Quênia","Кения","肯尼亚","肯尼亚",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"kg","قيرغيزستان","Kirgisistan","Kyrgyzstan","Kirguistán","Kirghizistan","キルギス","Quirguistão","Киргизия","吉尔吉斯斯坦","吉尔吉斯斯坦",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"kh","كمبوديا","Kambodscha","Cambodia","Camboya","Cambodge","カンボジア","Camboja","Камбоджа","柬埔寨","柬埔寨",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ki","كيريباتي","Kiribati","Kiribati","Kiribati","Kiribati","キリバス","Kiribati","Кирибати","基里巴斯","基里巴斯",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"km","جزر القمر","Komoren","Comoros","Comoras","Comores","コモロ","Comores","Коморы","科摩罗","科摩罗",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"kn","سانت كيتس ونيفيس","St. Kitts und Nevis","Saint Kitts and Nevis","San Cristóbal y Nieves","Saint-Christophe-et-Niévès","セントクリストファー・ネイビス","São Cristóvão e Nevis","Сент-Китс и Невис","圣基茨和尼维斯","圣基茨和尼维斯",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"kp","كوريا الشمالية","Nordkorea","North Korea","Corea del Norte","Corée du Nord","朝鮮民主主義人民共和国","Coreia do Norte","КНДР (Корейская Народно-Демократическая Республика)","朝鲜","朝鲜",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"kr","كوريا الجنوبية","Südkorea","South Korea","Corea del Sur","Corée du Sud","大韓民国","Coreia do Sul","Республика Корея","韩国","韩国",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"kw","الكويت","Kuwait","Kuwait","Kuwait","Koweït","クウェート","Kuwait","Кувейт","科威特","科威特",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ky","جزر كايمان","Kaimaninseln","Cayman Islands","Islas Caimán","Îles Caïmans","ケイマン諸島","Ilhas Cayman","Острова Кайман","开曼群岛","开曼群岛",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"kz","كازاخستان","Kasachstan","Kazakhstan","Kazajistán","Kazakhstan","カザフスタン","Cazaquistão","Казахстан","哈萨克斯坦","哈萨克斯坦",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"la","لاوس","Laos","Laos","Laos","Laos","ラオス人民民主共和国","Laos","Лаос","老挝","老挝",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"lb","لبنان","Libanon","Lebanon","Líbano","Liban","レバノン","Líbano","Ливан","黎巴嫩","黎巴嫩",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"lc","سانت لوسيا","St. Lucia","Saint Lucia","Santa Lucía","Sainte-Lucie","セントルシア","Santa Lúcia","Сент-Люсия","圣卢西亚","圣卢西亚",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"li","ليختنشتاين","Liechtenstein","Liechtenstein","Liechtenstein","Liechtenstein","リヒテンシュタイン","Liechtenstein","Лихтенштейн","列支敦斯登","列支敦斯登",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"lk","سريلانكا","Sri Lanka","Sri Lanka","Sri Lanka","Sri Lanka","スリランカ","Sri Lanka","Шри-Ланка","斯里蘭卡","斯里蘭卡",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"lr","ليبيريا","Liberia","Liberia","Liberia","Liberia","リベリア","Libéria","Либерия","利比里亚","利比里亚",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ls","ليسوتو","Lesotho","Lesotho","Lesoto","Lesotho","レソト","Lesoto","Лесото","賴索托","賴索托",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"lt","ليتوانيا","Litauen","Lithuania","Lituania","Lituanie","リトアニア","Lituânia","Литва","立陶宛","立陶宛",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"lu","لوكسمبورغ","Luxemburg","Luxembourg","Luxemburgo","Luxembourg","ルクセンブルク","Luxemburgo","Люксембург","卢森堡","卢森堡",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"lv","لاتفيا","Lettland","Latvia","Letonia","Lettonie","ラトビア","Letônia","Латвия","拉脫維亞","拉脫維亞",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ly","ليبيا","Libyen","Libya","Libia","Libye","リビア","Líbia","Ливия","利比亞","利比亞",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ma","المغرب","Marokko","Morocco","Marruecos","Maroc","モロッコ","Marrocos","Марокко","摩洛哥","摩洛哥",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mc","موناكو","Monaco","Monaco","Mónaco","Monaco","モナコ","Mônaco","Монако","摩納哥","摩納哥",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"md","مولدوفا","Moldawien","Moldova, Republic of","Moldavia","Moldavie","モルドバ共和国","Moldávia, República da","Молдавия","摩尔多瓦","摩尔多瓦",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"me","الجبل الأسود","Montenegro","Montenegro","Montenegro","Monténégro","モンテネグロ","Montenegro","Черногория","蒙特內哥羅","蒙特內哥羅",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mf","تجمع سان مارتين","Saint-Martin","Saint Martin (French part)","San Martín","Saint-Martin","サン・マルタン (フランス領)","São Martinho (parte francesa)","Сен-Мартен","法属圣马丁","法属圣马丁",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mg","مدغشقر","Madagaskar","Madagascar","Madagascar","Madagascar","マダガスカル","Madagáscar","Мадагаскар","马达加斯加","马达加斯加",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mh","جزر مارشال","Marshallinseln","Marshall Islands","Islas Marshall","Îles Marshall","マーシャル諸島","Ilhas Marshall","Маршалловы Острова","马绍尔群岛","马绍尔群岛",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mk","مقدونيا","Nordmazedonien","North Macedonia","Macedonia del Norte","Macédoine du Nord","北マケドニア","Macedônia do Norte","Северная Македония","北馬其頓","北馬其頓",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ml","مالي","Mali","Mali","Malí","Mali","マリ","Mali","Мали","马里","马里",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mm","ميانمار","Myanmar","Myanmar","Birmania","Birmanie","ミャンマー","Myanmar","Мьянма","緬甸","緬甸",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mn","منغوليا","Mongolei","Mongolia","Mongolia","Mongolie","モンゴル","Mongólia","Монголия","蒙古國","蒙古國",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mo","ماكاو","Macau","Macao","Macao","Macao","マカオ","Macau","Макао","澳門","澳門",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mp","جزر ماريانا الشمالية","Nördliche Marianen","Northern Mariana Islands","Islas Marianas del Norte","Îles Mariannes du Nord","北マリアナ諸島","Ilhas Marianas do Norte","Северные Марианские Острова","北馬里亞納群島","北馬里亞納群島",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mq","مارتينيك","Martinique","Martinique","Martinica","Martinique","マルティニーク","Martinica","Мартиника","马提尼克","马提尼克",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mr","موريتانيا","Mauretanien","Mauritania","Mauritania","Mauritanie","モーリタニア","Mauritânia","Мавритания","毛里塔尼亚","毛里塔尼亚",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ms","مونتسرات","Montserrat","Montserrat","Montserrat","Montserrat","モントセラト","Montserrat","Монтсеррат","蒙特塞拉特","蒙特塞拉特",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mt","مالطا","Malta","Malta","Malta","Malte","マルタ","Malta","Мальта","馬爾他","馬爾他",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mu","موريشيوس","Mauritius","Mauritius","Mauricio","Maurice","モーリシャス","Maurícia","Маврикий","模里西斯","模里西斯",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mv","جزر المالديف","Malediven","Maldives","Maldivas","Maldives","モルディブ","Maldivas","Мальдивы","馬爾地夫","馬爾地夫",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mw","مالاوي","Malawi","Malawi","Malaui","Malawi","マラウイ","Malawi","Малави","马拉维","马拉维",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mx","المكسيك","Mexiko","Mexico","México","Mexique","メキシコ","México","Мексика","墨西哥","墨西哥",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"my","ماليزيا","Malaysia","Malaysia","Malasia","Malaisie","マレーシア","Malásia","Малайзия","马来西亚","马来西亚",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"mz","موزمبيق","Mosambik","Mozambique","Mozambique","Mozambique","モザンビーク","Moçambique","Мозамбик","莫桑比克","莫桑比克",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"na","ناميبيا","Namibia","Namibia","Namibia","Namibie","ナミビア","Namíbia","Намибия","纳米比亚","纳米比亚",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"nc","كاليدونيا الجديدة","Neukaledonien","New Caledonia","Nueva Caledonia","Nouvelle-Calédonie","ニューカレドニア","Nova Caledônia","Новая Каледония","新喀里多尼亞","新喀里多尼亞",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ne","النيجر","Niger","Niger","Níger","Niger","ニジェール","Níger","Нигер","尼日尔","尼日尔",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"nf","جزيرة نورفولك","Norfolkinsel","Norfolk Island","Isla Norfolk","Île Norfolk","ノーフォーク島","Ilha Norfolk","Остров Норфолк","诺福克岛","诺福克岛",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ng","نيجيريا","Nigeria","Nigeria","Nigeria","Nigeria","ナイジェリア","Nigéria","Нигерия","奈及利亞","奈及利亞",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ni","نيكاراغوا","Nicaragua","Nicaragua","Nicaragua","Nicaragua","ニカラグア","Nicarágua","Никарагуа","尼加拉瓜","尼加拉瓜",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"nl","هولندا","Niederlande","Netherlands","Países Bajos","Pays-Bas","オランダ","Países Baixos","Нидерланды","荷蘭","荷蘭",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"no","النرويج","Norwegen","Norway","Noruega","Norvège","ノルウェー","Noruega","Норвегия","挪威","挪威",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"np","نيبال","Nepal","Nepal","Nepal","Népal","ネパール","Nepal","Непал","尼泊尔","尼泊尔",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"nr","ناورو","Nauru","Nauru","Nauru","Nauru","ナウル","Nauru","Науру","瑙鲁","瑙鲁",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"nu","نييوي","Niue","Niue","Niue","Niue","ニウエ","Niue","Ниуэ","纽埃","纽埃",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"nz","نيوزيلندا","Neuseeland","New Zealand","Nueva Zelanda","Nouvelle-Zélande","ニュージーランド","Nova Zelândia","Новая Зеландия","新西蘭","新西蘭",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"om","عمان","Oman","Oman","Omán","Oman","オマーン","Omã","Оман","阿曼","阿曼",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"pa","بنما","Panama","Panama","Panamá","Panama","パナマ","Panamá","Панама","巴拿马","巴拿马",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"pe","بيرو","Peru","Peru","Perú","Pérou","ペルー","Peru","Перу","秘魯","秘魯",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"pf","بولينزيا الفرنسية","Französisch-Polynesien","French Polynesia","Polinesia Francesa","Polynésie française","フランス領ポリネシア","Polinésia Francesa","Французская Полинезия","法屬玻里尼西亞","法屬玻里尼西亞",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"pg","بابوا غينيا الجديدة","Papua-Neuguinea","Papua New Guinea","Papúa Nueva Guinea","Papouasie-Nouvelle-Guinée","パプアニューギニア","Papua Nova Guiné","Папуа — Новая Гвинея","巴布亚新几内亚","巴布亚新几内亚",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ph","الفلبين","Philippinen","Philippines","Filipinas","Philippines","フィリピン","Filipinos","Филиппины","菲律賓","菲律賓",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"pk","باكستان","Pakistan","Pakistan","Pakistán","Pakistan","パキスタン","Paquistão","Пакистан","巴基斯坦","巴基斯坦",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"pl","بولندا","Polen","Poland","Polonia","Pologne","ポーランド","Polônia","Польша","波蘭","波蘭",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"pm","سان بيير وميكلون","Saint-Pierre und Miquelon","Saint Pierre and Miquelon","San Pedro y Miquelón","Saint-Pierre-et-Miquelon","サンピエール島・ミクロン島","São Pedro e Miquelon","Сен-Пьер и Микелон","圣皮埃尔和密克隆","圣皮埃尔和密克隆",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"pn","جزر بيتكيرن","Pitcairninseln","Pitcairn","Islas Pitcairn","Îles Pitcairn","ピトケアン","Pitcairn","Острова Питкэрн","皮特凯恩群岛","皮特凯恩群岛",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"pr","بورتوريكو","Puerto Rico","Puerto Rico","Puerto Rico","Porto Rico","プエルトリコ","Porto Rico","Пуэрто-Рико","波多黎各","波多黎各",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ps","فلسطين","Staat Palästina","Palestine, State of","Palestina","Palestine","パレスチナ","Palestina","Государство Палестина","巴勒斯坦","巴勒斯坦",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"pt","البرتغال","Portugal","Portugal","Portugal","Portugal","ポルトガル","Portugal","Португалия","葡萄牙","葡萄牙",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"pw","بالاو","Palau","Palau","Palaos","Palaos","パラオ","Palau","Палау","帛琉","帛琉",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"py","باراغواي","Paraguay","Paraguay","Paraguay","Paraguay","パラグアイ","Paraguai","Парагвай","巴拉圭","巴拉圭",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"qa","قطر","Katar","Qatar","Catar","Qatar","カタール","Catar","Катар","卡塔尔","卡塔尔",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"re","لا ريونيون","Réunion","Reunion","Reunión","La Réunion","レユニオン","Reunião","Реюньон","留尼汪","留尼汪",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ro","رومانيا","Rumänien","Romania","Rumania","Roumanie","ルーマニア","Romênia","Румыния","羅馬尼亞","羅馬尼亞",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"rs","صربيا","Serbien","Serbia","Serbia","Serbie","セルビア","Sérvio","Сербия","塞爾維亞","塞爾維亞",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ru","روسيا","Russische Föderation","Russia","Rusia","Russie","ロシア","Rússia","Россия","俄羅斯","俄羅斯",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"rw","رواندا","Ruanda","Rwanda","Ruanda","Rwanda","ルワンダ","Ruanda","Руанда","卢旺达","卢旺达",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sa","السعودية","Saudi-Arabien","Saudi Arabia","Arabia Saudita","Arabie saoudite","サウジアラビア","Arábia Saudita","Саудовская Аравия","沙烏地阿拉伯","沙烏地阿拉伯",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sb","جزر سليمان","Salomonen","Solomon Islands","Islas Salomón","Salomon","ソロモン諸島","Ilhas Salomão","Соломоновы Острова","所罗门群岛","所罗门群岛",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sc","سيشل","Seychellen","Seychelles","Seychelles","Seychelles","セーシェル","Seychelles","Сейшельские Острова","塞舌尔","塞舌尔",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sd","السودان","Sudan","Sudan","Sudán","Soudan","スーダン","Sudão","Судан","苏丹","苏丹",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"se","السويد","Schweden","Sweden","Suecia","Suède","スウェーデン","Suécia","Швеция","瑞典","瑞典",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sg","سنغافورة","Singapur","Singapore","Singapur","Singapour","シンガポール","Cingapura","Сингапур","新加坡","新加坡",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sh","سانت هيلانة وأسينشين وتريستان دا كونا","St. Helena","Saint Helena, Ascension and Tristan da Cunha","Santa Elena, Ascensión y Tristán de Acuña","Sainte-Hélène, Ascension et Tristan da Cunha","セントヘレナ・アセンションおよびトリスタンダクーニャ","Santa Helena, Ascensão e Tristão da Cunha","Острова Святой Елены, Вознесения и Тристан-да-Кунья","圣赫勒拿、阿森松和特里斯坦-达库尼亚","圣赫勒拿、阿森松和特里斯坦-达库尼亚",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"si","سلوفينيا","Slowenien","Slovenia","Eslovenia","Slovénie","スロベニア","Eslovênia","Словения","斯洛維尼亞","斯洛維尼亞",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sj","سفالبارد ويان ماين","Svalbard und Jan Mayen","Svalbard and Jan Mayen","Svalbard y Jan Mayen","Svalbard et ile Jan Mayen","スヴァールバル諸島およびヤンマイエン島","Svalbard e Jan Mayen","Шпицберген и Ян-Майен","斯瓦尔巴和扬马延","斯瓦尔巴和扬马延",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sk","سلوفاكيا","Slowakei","Slovakia","Eslovaquia","Slovaquie","スロバキア","Eslováquia","Словакия","斯洛伐克","斯洛伐克",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sl","سيراليون","Sierra Leone","Sierra Leone","Sierra Leona","Sierra Leone","シエラレオネ","Serra Leoa","Сьерра-Леоне","塞拉利昂","塞拉利昂",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sm","سان مارينو","San Marino","San Marino","San Marino","Saint-Marin","サンマリノ","San Marino","Сан-Марино","圣马力诺","圣马力诺",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sn","السنغال","Senegal","Senegal","Senegal","Sénégal","セネガル","Senegal","Сенегал","塞内加尔","塞内加尔",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"so","الصومال","Somalia","Somalia","Somalia","Somalie","ソマリア","Somália","Сомали","索馬利亞","索馬利亞",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sr","سورينام","Suriname","Suriname","Surinam","Suriname","スリナム","Suriname","Суринам","苏里南","苏里南",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ss","جنوب السودان","Südsudan","South Sudan","Sudán del Sur","Soudan du Sud","南スーダン","Sudão do Sul","Южный Судан","南蘇丹","南蘇丹",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"st","ساو تومي وبرينسيب","São Tomé und Príncipe","Sao Tome and Principe","Santo Tomé y Príncipe","Sao Tomé-et-Principe","サントメ・プリンシペ","São Tomé e Príncipe","Сан-Томе и Принсипи","聖多美和普林西比","聖多美和普林西比",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sv","السلفادور","El Salvador","El Salvador","El Salvador","Salvador","エルサルバドル","El Salvador","Сальвадор","薩爾瓦多","薩爾瓦多",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sx","سينت مارتن","Sint Maarten","Sint Maarten (Dutch part)","San Martín","Saint-Martin","シント・マールテン (オランダ領)","São Martinho (parte holandesa)","Синт-Мартен","聖馬丁","聖馬丁",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sy","سوريا","Syrien","Syria","Siria","Syrie","シリア・アラブ共和国","Sírio","Сирия","叙利亚","叙利亚",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"sz","سوازيلاند","Swasiland","Eswatini","Suazilandia","Swaziland","エスワティニ","Eswatini","Эсватини","斯威士兰","斯威士兰",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tc","جزر توركس وكايكوس","Turks- und Caicosinseln","Turks and Caicos Islands","Islas Turcas y Caicos","Îles Turques-et-Caïques","タークス・カイコス諸島","Ilhas Turks e Caicos","Теркс и Кайкос","特克斯和凯科斯群岛","特克斯和凯科斯群岛",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"td","تشاد","Tschad","Chad","Chad","Tchad","チャド","Chade","Чад","乍得","乍得",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tf","أراض فرنسية جنوبية وأنتارتيكية","Französische Süd- und Antarktisgebiete","French Southern Territories","Tierras Australes y Antárticas Francesas","Terres australes et antarctiques françaises","フランス領南方・南極地域","Territórios Franceses do Sul","Французские Южные и Антарктические Территории","法属南方和南极洲领地","法属南方和南极洲领地",7,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tg","توغو","Togo","Togo","Togo","Togo","トーゴ","Ir","Того","多哥","多哥",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"th","تايلاند","Thailand","Thailand","Tailandia","Thaïlande","タイ","Tailândia","Таиланд","泰國","泰國",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tj","طاجيكستان","Tadschikistan","Tajikistan","Tayikistán","Tadjikistan","タジキスタン","Tajiquistão","Таджикистан","塔吉克斯坦","塔吉克斯坦",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tk","توكيلاو","Tokelau","Tokelau","Tokelau","Tokelau","トケラウ","Tokelau","Токелау","托克勞","托克勞",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tl","تيمور الشرقية","Timor-Leste","Timor-Leste","Timor Oriental","Timor oriental","東ティモール","Timor-Leste","Восточный Тимор","东帝汶","东帝汶",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tm","تركمانستان","Turkmenistan","Turkmenistan","Turkmenistán","Turkménistan","トルクメニスタン","Turquemenistão","Туркмения","土库曼斯坦","土库曼斯坦",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tn","تونس","Tunesien","Tunisia","Túnez","Tunisie","チュニジア","Tunísia","Тунис","突尼西亞","突尼西亞",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"to","تونغا","Tonga","Tonga","Tonga","Tonga","トンガ","Tonga","Тонга","汤加","汤加",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tr","تركيا","Türkei","Turkey","Turquía","Turquie","トルコ","Turquia","Турция","土耳其","土耳其",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tt","ترينيداد وتوباغو","Trinidad und Tobago","Trinidad and Tobago","Trinidad y Tobago","Trinité-et-Tobago","トリニダード・トバゴ","Trindade e Tobago","Тринидад и Тобаго","千里達及托巴哥","千里達及托巴哥",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tv","توفالو","Tuvalu","Tuvalu","Tuvalu","Tuvalu","ツバル","Tuvalu","Тувалу","图瓦卢","图瓦卢",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tw","تايوان","Taiwan","Taiwan","Taiwán","Taïwan","台湾","Taiwan","Китайская Республика","台湾","台湾",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"tz","تنزانيا","Tansania, Vereinigte Republik","Tanzania, United Republic of","Tanzania","Tanzanie","タンザニア","Tanzânia, República Unida da","Танзания","坦桑尼亚","坦桑尼亚",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ua","أوكرانيا","Ukraine","Ukraine","Ucrania","Ukraine","ウクライナ","Ucrânia","Украина","烏克蘭","烏克蘭",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ug","أوغندا","Uganda","Uganda","Uganda","Ouganda","ウガンダ","Uganda","Уганда","乌干达","乌干达",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"uk","المملكة المتحدة","Vereinigtes Königreich Großbritannien und Nordirland","United Kingdom","Reino Unido","Royaume-Uni","イギリス","Reino Unido","Великобритания","英國","英國",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"um","جزر الولايات المتحدة الصغيرة النائية","United States Minor Outlying Islands","United States Minor Outlying Islands","Islas ultramarinas de Estados Unidos","Îles mineures éloignées des États-Unis","合衆国領有小離島","Ilhas Menores Distantes dos Estados Unidos","Внешние малые острова (США)","美國本土外小島嶼","美國本土外小島嶼",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"us","الولايات المتحدة","Vereinigte Staaten von Amerika","United States of America","Estados Unidos","États-Unis","アメリカ合衆国","Estados Unidos da America","США","美國","美國",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"uy","الأوروغواي","Uruguay","Uruguay","Uruguay","Uruguay","ウルグアイ","Uruguai","Уругвай","乌拉圭","乌拉圭",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"uz","أوزبكستان","Usbekistan","Uzbekistan","Uzbekistán","Ouzbékistan","ウズベキスタン","Usbequistão","Узбекистан","乌兹别克斯坦","乌兹别克斯坦",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"va","الفاتيكان","Vatikanstadt","Holy See","Ciudad del Vaticano","Saint-Siège","バチカン市国","Cidade do Vaticano","Ватикан","梵蒂冈","梵蒂冈",3,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"vc","سانت فينسنت والغرينادين","St. Vincent und die Grenadinen","Saint Vincent and the Grenadines","San Vicente y las Granadinas","Saint-Vincent-et-les-Grenadines","セントビンセントおよびグレナディーン諸島","São Vicente e Granadinas","Сент-Винсент и Гренадины","圣文森特和格林纳丁斯","圣文森特和格林纳丁斯",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ve","فنزويلا","Venezuela","Venezuela (Bolivarian Republic of)","Venezuela","Venezuela","ベネズエラ・ボリバル共和国","Venezuela","Венесуэла","委內瑞拉","委內瑞拉",5,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"vg","جزر العذراء البريطانية","Britische Jungferninseln","Virgin Islands (British)","Islas Vírgenes Británicas","Îles Vierges britanniques","イギリス領ヴァージン諸島","Ilhas Virgens Britânicas","Виргинские Острова (Великобритания)","英屬維爾京群島","英屬維爾京群島",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"vi","جزر العذراء الأمريكية","Amerikanische Jungferninseln","Virgin Islands (U.S.)","Islas Vírgenes de los Estados Unidos","Îles Vierges des États-Unis","アメリカ領ヴァージン諸島","Ilhas Virgens (EUA)","Виргинские Острова (США)","美屬維爾京群島","美屬維爾京群島",4,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"vn","فيتنام","Vietnam","Vietnam","Vietnam","Viêtnam","ベトナム","Vietnã","Вьетнам","越南","越南",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"vu","فانواتو","Vanuatu","Vanuatu","Vanuatu","Vanuatu","バヌアツ","Vanuatu","Вануату","瓦努阿圖","瓦努阿圖",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"wf","والس وفوتونا","Wallis und Futuna","Wallis and Futuna","Wallis y Futuna","Wallis-et-Futuna","ウォリス・フツナ","Wallis e Futuna","Уоллис и Футуна","瓦利斯和富圖納","瓦利斯和富圖納",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ws","ساموا","Samoa","Samoa","Samoa","Samoa","サモア","Samoa","Самоа","萨摩亚","萨摩亚",6,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"ye","اليمن","Jemen","Yemen","Yemen","Yémen","イエメン","Iémen","Йемен","葉門","葉門",2,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"yt","مايوت","Mayotte","Mayotte","Mayotte","Mayotte","マヨット","Mayotte","Майотта","马约特","马约特",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"za","جنوب أفريقيا","Südafrika","South Africa","Sudáfrica","Afrique du Sud","南アフリカ","África do Sul","ЮАР","南非","南非",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"zm","زامبيا","Sambia","Zambia","Zambia","Zambie","ザンビア","Zâmbia","Замбия","尚比亞","尚比亞",1,VAL_STATUS_ACTIVE}).
            AppendValues([]interface{}{"zw","زيمبابوي","Simbabwe","Zimbabwe","Zimbabue","Zimbabwe","ジンバブエ","Zimbábue","Зимбабве","辛巴威","辛巴威",1,VAL_STATUS_ACTIVE})

    return c.BaseClient.Run()
}
