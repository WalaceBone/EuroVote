package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type RollCallVote struct {
	XMLName xml.Name `xml:"RollCallVote"`
	Result  Result   `xml:"Result"`
}

type Result struct {
	Identifier  string      `xml:"Identifier,attr"`
	DlvId       string      `xml:"DlvId,attr"`
	Date        string      `xml:"Date,attr"`
	Description Description `xml:"Description"`
	For         For         `xml:"For"`
	Against     Against     `xml:"Against"`
	Abstention  Abstention  `xml:"Abstention"`
	Intentions  Intentions  `xml:"Intentions"`
}

type Description struct {
	Text string `xml:"Text"`
}

type For struct {
	Number             string               `xml:"Number,attr"`
	PoliticalGroupList []PoliticalGroupList `xml:"PoliticalGroup.List"`
}

type Against struct {
	Number             string               `xml:"Number,attr"`
	PoliticalGroupList []PoliticalGroupList `xml:"PoliticalGroup.List"`
}

type Abstention struct {
	Number             string               `xml:"Number,attr"`
	PoliticalGroupList []PoliticalGroupList `xml:"PoliticalGroup.List"`
}

type Intentions struct {
	Number             string               `xml:"Number,attr"`
	PoliticalGroupList []PoliticalGroupList `xml:"PoliticalGroup.List"`
}

type PoliticalGroupList struct {
	Identifier string                 `xml:"Identifier,attr"`
	Members    []PoliticalGroupMember `xml:"PoliticalGroup.Member"`
}

type PoliticalGroupMember struct {
	MepId  string `xml:"MepId,attr"`
	PersId string `xml:"PersId,attr"`
	Name   string `xml:",chardata"` // Represents the text content of the element
	// Add more fields as needed to represent other attributes or elements within PoliticalGroup.Member
}

func main() {
	// Sample XML data (replace this with your actual XML data)
	xmlData := `
<RollCallVote.Result Identifier="157790" DlvId="939445" Date="2023-09-12 12:08:42">
    <RollCallVote.Description.Text> Étiquetage des aliments biologiques pour animaux familiers -
        Labelling of organic pet food - Kennzeichnung von ökologischem/biologischem Heimtierfutter - <a
            href="#reds:iPlRp/A-9-2023-0159" data-rel="reds" redmap-uri="/reds:iPlRp/A-9-2023-0159">
        A9-0159/2023</a> - Martin Häusling - Accord provisoire - Am 17 </RollCallVote.Description.Text>
    <Result.For Number="592">
        <Result.PoliticalGroup.List Identifier="ECR">
            <PoliticalGroup.Member.Name MepId="7244" PersId="198096">Aguilar</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4978" PersId="4746">Berlato</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5112" PersId="23788">Bielan</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6864" PersId="197467">Bourgeois</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7214" PersId="197829">Buxadé Villalba</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5490" PersId="28372">Czarnecki</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6511" PersId="124873">Dzhambazki</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5472" PersId="28353">Fotyga</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7250" PersId="198490">Fragkos</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7179" PersId="197794">Gemma</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6989" PersId="197596">Hakkarainen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7289" PersId="218349">Hoogeveen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7290" PersId="221463">Ilčić</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6909" PersId="197516">Jaki</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6911" PersId="197518">Jurgiel</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7151" PersId="197767">Jurzyca</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6866" PersId="197469">Kanko</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6525" PersId="124887">Karski</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6912" PersId="197519">Kempa</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6913" PersId="197520">Kloc</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6923" PersId="197530">Kopcińska</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6530" PersId="124891">Krasnodębski</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6925" PersId="197532">Kruk</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5507" PersId="28389">Kuźmiuk</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5969" PersId="96796">Legutko</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6608" PersId="124996">Lundgren</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6931" PersId="197538">Mazurek</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7209" PersId="197824">Milazzo</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6934" PersId="197541">Możdżanowska</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7270" PersId="204414">de la Pisa Carrión</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6937" PersId="197544">Rafalska</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7160" PersId="197776">Rooken</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7097" PersId="197709">Roos</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7157" PersId="197773">Ruissen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6718" PersId="132366">Ruohonen-Lerner</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6938" PersId="197545">Rzońca</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5416" PersId="28297">Saryusz-Wolski</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7210" PersId="197825">Stancanelli</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6946" PersId="197553">Szydło</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7265" PersId="204346">Tarczyński</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7044" PersId="197655">Terheş</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7216" PersId="197831">Tertsch</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6952" PersId="197559">Tobiszowski</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5878" PersId="96697">Tomaszewski</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5892" PersId="96713">Tošenovský</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6683" PersId="125106">Van Overtveldt</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6930" PersId="197537">Vondra</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6945" PersId="197552">Vrecionová</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6956" PersId="197566">Waszczykowski</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6515" PersId="124877">Wiśniewska</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5052" PersId="23712">Zahradil</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6962" PersId="197572">Zalewska</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5598" PersId="28615">Zīle</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6522" PersId="124884">Złotowski</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="ID">
            <PoliticalGroup.Member.Name MepId="6872" PersId="197475">Anderson</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7079" PersId="197691">Androuët</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6593" PersId="124973">Annemans</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6988" PersId="131580">Bardella</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7001" PersId="197610">Basso</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6874" PersId="132191">Beck</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7068" PersId="197680">Beigneux</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6405" PersId="124771">Bilde</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7172" PersId="197788">Bonfrisco</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7000" PersId="101039">Borchia</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7076" PersId="182995">Bruna</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6876" PersId="128483">Buchheit</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6981" PersId="35016">Campomenosi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7181" PersId="197796">Casanova</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7170" PersId="197786">Ceccardi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7302" PersId="236052">Chagnon</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7007" PersId="197616">Conte</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6999" PersId="197608">Da Re</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7300" PersId="236050">Dauchy</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6869" PersId="197472">De Man</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6972" PersId="197582">Gancia</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7018" PersId="197628">Garraud</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7311" PersId="239255">Gazzini</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7312" PersId="239256">Ghidoni</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7184" PersId="197799">Grant</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7019" PersId="94649">Griset</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7246" PersId="198176">Haider</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6751" PersId="189065">Jamet</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7017" PersId="197627">Joron</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6879" PersId="197482">Kuhs</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7275" PersId="204421">Lacapelle</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6771" PersId="192635">Lancini</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6372" PersId="124738">Lebreton</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6880" PersId="197483">Limmer</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7002" PersId="197611">Lizzi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6886" PersId="197493">Madison</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7013" PersId="197623">Mariani</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6627" PersId="38511">Mayer</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7301" PersId="236051">Minardi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7075" PersId="197687">Olivier</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6979" PersId="99283">Panza</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6881" PersId="197488">Reil</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7171" PersId="197787">Rinaldi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7322" PersId="243912">Rossi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7078" PersId="197690">Rougé</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6968" PersId="197578">Sardone</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7191" PersId="197806">Tardino</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6970" PersId="197580">Tovaglieri</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6610" PersId="125001">Vilimsky</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6513" PersId="124875">Vistisen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6976" PersId="197586">Zambelli</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6414" PersId="124780">Zanni</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6882" PersId="197489">Zimniok</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="NI">
            <PoliticalGroup.Member.Name MepId="6394" PersId="124760">Bay</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6411" PersId="124777">Beghin</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6343" PersId="124712">Bocskor</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6446" PersId="124812">Castaldo</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6051" PersId="96880">Cozzolino</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7314" PersId="239258">Danzì</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6345" PersId="124714">Deli</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7207" PersId="197822">Donato</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6470" PersId="124833">Ferrara</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7204" PersId="197819">Georgoulis</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7243" PersId="198063">Gyöngyösi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6002" PersId="96830">Győri</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7328" PersId="247709">Haga</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6318" PersId="124586">Hidvéghi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7016" PersId="197626">Juvin</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6686" PersId="125109">Kaili</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6836" PersId="197438">Kolakušić</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7123" PersId="197740">Konstantinou</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6001" PersId="96829">Kósa</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6759" PersId="190518">Meuthen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7010" PersId="197619">Pignedoli</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7020" PersId="197629">Pirbakas</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7276" PersId="204443">Ponsatí Obiols</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7257" PersId="202351">Puigdemont i Casamajó</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7286" PersId="213330">Schaller-Baross</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5624" PersId="29579">Tarabella</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6965" PersId="27714">Tóth</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5879" PersId="96698">Uspaskich</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="PPE">
            <PoliticalGroup.Member.Name MepId="6883" PersId="197490">Adamowicz</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6755" PersId="189525">Ademov</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6468" PersId="124831">Adinolfi Isabella</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7219" PersId="197836">Alexandrov Yordanov</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5980" PersId="96808">Arias Echeverría</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6885" PersId="197492">Arłukowicz</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7083" PersId="197695">Asimakopoulou</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7060" PersId="197671">Băsescu</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6927" PersId="197534">Bellamy</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7067" PersId="197679">Benjumea Benjumea</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6807" PersId="197408">Bentele</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7162" PersId="197778">Berendsen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6809" PersId="197410">Berger</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7037" PersId="197648">Bernhuber</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7155" PersId="197771">Bilčík</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7026" PersId="197637">Blaga</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7025" PersId="197636">Bogdan</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6613" PersId="125004">Bogovič</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7291" PersId="226260">Braunsberger-Reinhold</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6617" PersId="125012">Buda</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5729" PersId="38420">Buşoi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6039" PersId="96867">Carvalho</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5252" PersId="28122">Casa</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5341" PersId="28219">Caspary</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5508" PersId="28390">del Castillo Vera</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6499" PersId="124861">Chinnici</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6600" PersId="124988">Clune</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="3800" PersId="1892">Coelho</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6929" PersId="197536">Colin-Oesterlé</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7330" PersId="247737">Collado Jiménez</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5948" PersId="96775">Comi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5921" PersId="96747">Danjean</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7261" PersId="204333">De Meo</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6762" PersId="190774">Didier</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6870" PersId="197473">Doleschal</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5960" PersId="96787">Dorfmann</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6903" PersId="197510">Duda</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6810" PersId="99945">Düpont</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5348" PersId="28226">Ehler</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5983" PersId="96811">Estaràs Ferragut</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6928" PersId="197535">Evren</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7038" PersId="197649">Falcă</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="3833" PersId="1917">Ferber</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6068" PersId="96899">Fernandes</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7108" PersId="197720">Fitzgerald</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6813" PersId="197414">Fourlas</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6906" PersId="197513">Frankowski</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6852" PersId="197455">Franssen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4659" PersId="2341">Gahler</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6834" PersId="197436">Geuking</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6441" PersId="124807">Gieseke</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7253" PersId="202036">Glavak</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6907" PersId="197514">Halicki</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7035" PersId="197646">Hava</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6811" PersId="197412">Herbst</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6534" PersId="124895">Hetman</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4987" PersId="5565">Hortefeux</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5952" PersId="96779">Hübner</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5945" PersId="96772">Jahr</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6910" PersId="197517">Jarubas</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7223" PersId="197840">Juknevičienė</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5944" PersId="96771">Kalinowski</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6095" PersId="96934">Kalniete</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7222" PersId="197839">Kanev</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4669" PersId="4246">Karas</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6666" PersId="125068">Kefalogiannis</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5855" PersId="96668">Kelly</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6805" PersId="197406">Kokalari</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6196" PersId="97968">Kovatchev</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7226" PersId="197843">Kubilius</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5727" PersId="38398">de Lange</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6793" PersId="197393">Lega</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6592" PersId="95074">Lenaers</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7264" PersId="204336">Lexmann</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="3855" PersId="1927">Liese</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6442" PersId="124808">Lins</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5517" PersId="28399">López-Istúriz White</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5964" PersId="96791">Łukacijewska</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6855" PersId="197458">Lutgen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6440" PersId="124806">McAllister</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7218" PersId="197835">Maldeikienė</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4964" PersId="4560">Manders</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6761" PersId="190713">Mandl</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5646" PersId="33982">Marinescu</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7284" PersId="209896">Markey</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6465" PersId="124828">Martusciello</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6096" PersId="96936">Mato</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6483" PersId="98341">Maydell</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7224" PersId="197841">Mažylis</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7051" PersId="197662">Meimarakis</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7173" PersId="130256">Melbārde</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6121" PersId="96978">Melo</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5518" PersId="28400">Millán Mon</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6366" PersId="124734">Monteiro de Aguiar</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7099" PersId="197711">Montserrat</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6368" PersId="72779">Morano</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6825" PersId="197427">Mortler</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7036" PersId="197647">Motreanu</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6436" PersId="124802">Mureşan</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5547" PersId="28429">Mussolini</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4712" PersId="4289">Niebler</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6332" PersId="124701">Niedermayer</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7255" PersId="202112">Nistor</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5413" PersId="28294">Novak</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6935" PersId="197542">Ochojska</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5702" PersId="36392">Patriciello</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7323" PersId="243979">Peppucci</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7120" PersId="197738">Pereira Lídia</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5346" PersId="28224">Pieper</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5788" PersId="40599">Pietikäinen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6335" PersId="124704">Polčák</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6803" PersId="197404">Polfjärd</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7149" PersId="197765">Pollák</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6700" PersId="125706">Pospíšil</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6488" PersId="124850">Radev</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6750" PersId="188945">Radtke</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6072" PersId="96903">Rangel</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6812" PersId="197413">Ressler</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7266" PersId="204368">Sagartz</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6698" PersId="125670">Salini</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6370" PersId="24594">Sander</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6259" PersId="112611">Sarvamaa</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7045" PersId="197656">Schmiedtbauer</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6821" PersId="197422">Schneider</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5345" PersId="28223">Schwab</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6823" PersId="197425">Seekatz</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6824" PersId="197426">Simon</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6790" PersId="197390">Skyttedal</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6341" PersId="124710">Šojdrová</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6816" PersId="197417">Sokol</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6663" PersId="125064">Spyraki</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7317" PersId="239271">Stavrou</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6547" PersId="124929">Štefanec</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7277" PersId="204449">Terras</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7056" PersId="197667">Thaler</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6801" PersId="197402">Tobé</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7061" PersId="197672">Tomac</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6681" PersId="125104">Tomc</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6706" PersId="129164">Vandenkendelaere</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5930" PersId="96756">Verheyen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7062" PersId="98582">Vincze</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6356" PersId="124726">Virkkunen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5934" PersId="96761">Voss</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6664" PersId="125065">Vozemberg-Vrionidi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7183" PersId="197798">Vuolo</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7233" PersId="197863">Walsh</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6827" PersId="197429">Walsmann</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6804" PersId="197405">Warborn</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6908" PersId="197515">Weiss</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4643" PersId="2323">Wieland</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5782" PersId="39725">Winkler</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7041" PersId="197652">Winzig</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6820" PersId="197421">Wiseler-Lima</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6665" PersId="125067">Zagorakis</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6344" PersId="124713">Zdechovský</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7012" PersId="197621">Zoido Álvarez</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6734" PersId="185341">Zovko</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6094" PersId="96933">Zver</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="Renew">
            <PoliticalGroup.Member.Name MepId="6799" PersId="197400">Al-Sahlani</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7327" PersId="247122">Amalric</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7260" PersId="204332">Andrews</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6327" PersId="124696">Ansip</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6400" PersId="124766">Auštrevičius</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7165" PersId="197781">Azmani</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7163" PersId="197779">Bauzá Díaz</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6944" PersId="197551">Bijoux</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6089" PersId="96922">Bilbao Barandica</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7057" PersId="197668">Botoş</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6967" PersId="197577">Boyer</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6966" PersId="197576">Brunet</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6973" PersId="126644">Cañas</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5891" PersId="96711">Canfin</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6899" PersId="197505">Chabaud</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6339" PersId="124708">Charanzová</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6860" PersId="197463">Chastel</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6951" PersId="197558">Christensen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7271" PersId="204416">Cicurel</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7021" PersId="197631">Cioloş</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6978" PersId="197588">Cseh</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6458" PersId="124821">Danti</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6898" PersId="197504">Decerle</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6340" PersId="124709">Dlabajová</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6987" PersId="197595">Donáth</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6980" PersId="197589">Farreng</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6770" PersId="192634">Ferrandino</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6818" PersId="197419">Flego</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7283" PersId="209140">Gheorghe</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6840" PersId="197443">Glueck</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="1793" PersId="840">Goerens</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7273" PersId="204419">Gozi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6849" PersId="197452">Grošelj</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6950" PersId="197557">Grudler</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6936" PersId="197543">Guetta</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6841" PersId="197444">Hahn Svenja</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6959" PersId="135511">Hayer</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6919" PersId="197526">Hlaváček</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7154" PersId="197770">Hojsík</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6622" PersId="58789">Huitema</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7014" PersId="197624">Ijabs</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5385" PersId="28266">in 't Veld</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6844" PersId="197447">Joveva</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6975" PersId="197585">Karleskind</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6800" PersId="197401">Karlsbro</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6766" PersId="191693">Katainen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7203" PersId="197818">Kelleher</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5021" PersId="22858">Keller Fabienne</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6921" PersId="197528">Knotek</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6842" PersId="197445">Körner</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6917" PersId="118949">Kovařík</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6662" PersId="125063">Kyrtsos</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6504" PersId="124866">Kyuchyuk</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6887" PersId="197494">Loiseau</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5889" PersId="96709">Løkkegaard</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7319" PersId="239973">Rasmussen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6957" PersId="197567">Melchior</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6690" PersId="125128">Mihaylova</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7285" PersId="212855">Mituța</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6500" PersId="124862">Müller</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6760" PersId="190519">Nagtegaal</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6614" PersId="125005">Nart</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6830" PersId="197432">Oetjen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7299" PersId="234344">Orville</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6644" PersId="125038">Pagazaurtundúa</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6953" PersId="197563">Pekkarinen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6510" PersId="124872">Petersen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7052" PersId="197663">Pîslaru</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7306" PersId="237320">Poptcheva</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7318" PersId="239972">Poulsen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7235" PersId="197868">Rafaela</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4676" PersId="4253">Ries</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6056" PersId="96885">Riquet</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4767" PersId="4344">Rodríguez Ramos</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6902" PersId="197508">Séjourné</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6817" PersId="197418">Semedo</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7168" PersId="197784">Solís Pérez</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7058" PersId="134605">Strugariu</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5949" PersId="96776">Thun und Hohenstein</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6940" PersId="197547">Tolleret</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6331" PersId="124700">Toom</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6281" PersId="114268">Torvalds</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7054" PersId="197665">Tudorache</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6710" PersId="130100">Vautmans</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6896" PersId="197502">Vedrenne</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6159" PersId="97058">Verhofstadt</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7268" PersId="204400">Vázquez Lázara</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7287" PersId="214839">Wiesner</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7153" PersId="197769">Wiezik</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6904" PersId="197511">Yenbou</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6971" PersId="197581">Yon-Courtin</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6695" PersId="125237">Zullo</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="S&D">
            <PoliticalGroup.Member.Name MepId="6802" PersId="197403">Agius Saliba</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6651" PersId="125045">Aguilera</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7305" PersId="237224">Albuquerque</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7167" PersId="197783">Ameriks</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7254" PersId="202073">Angel</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6977" PersId="197587">Ara-Kovács</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7040" PersId="197651">Avram</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7329" PersId="247735">Ballarín Cereza</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6888" PersId="197495">Balt</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7193" PersId="197808">Bartolo</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6889" PersId="197496">Belka</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6505" PersId="124867">Benifei</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6796" PersId="197396">Bergkvist</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6891" PersId="197498">Biedroń</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6833" PersId="197435">Bischoff</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5867" PersId="96681">Blinkevičiūtė</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6264" PersId="112748">Borzan</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5465" PersId="28346">Bresso</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6843" PersId="197446">Brglez</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4690" PersId="4267">Bullmann</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6838" PersId="197440">Burkhardt</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7252" PersId="199996">Carvalhais</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7030" PersId="197641">Cerdas</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7166" PersId="197782">Chahim</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6901" PersId="197507">Cimoszewicz</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7325" PersId="245018">Clergeau</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7321" PersId="240478">Covassi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5661" PersId="33997">Crețu</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6794" PersId="197394">Cutajar</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7293" PersId="228604">De Basso</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6062" PersId="96891">De Castro</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6969" PersId="197579">Dobrev</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7248" PersId="198329">Durá Ferrandis</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6324" PersId="124693">Durand</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7307" PersId="237465">Ecke</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7282" PersId="209091">Engerer</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6652" PersId="125046">Fernández</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6791" PersId="197391">Fritzon</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6955" PersId="101585">Fuglsang</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7090" PersId="197702">Gálvez Muñoz</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7105" PersId="197717">García Del Blanco</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7125" PersId="197742">García Muñoz</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5417" PersId="28298">García Pérez</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6130" PersId="96991">Gardiazabal Rubial</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6005" PersId="96833">Geier</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7082" PersId="197694">Glucksmann</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7111" PersId="197728">González</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7106" PersId="197718">González Casares</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6419" PersId="124785">Grapini</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7009" PersId="197618">Gualmini</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6107" PersId="96952">Guillaume</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7146" PersId="34578">Hajšel</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7070" PersId="197682">Heide</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7185" PersId="197800">Heinäluoma</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7122" PersId="122978">Homs Ginel</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7229" PersId="197846">Hristov</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6263" PersId="112747">Jerković</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6624" PersId="125021">Jongerius</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6884" PersId="197491">Kaljurand</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6916" PersId="197523">Kohut</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6459" PersId="124822">Köster</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7085" PersId="197697">Lalucq</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7086" PersId="197698">Larrouturou</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7297" PersId="230085">Laureti</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7024" PersId="197635">Leitão-Marques</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5094" PersId="23768">Liberadzki</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6648" PersId="125042">López</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5984" PersId="96812">López Aguilar</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7292" PersId="228286">Lucke</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7109" PersId="197721">Luena</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7107" PersId="197719">Maestre Martín De Almagro</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7027" PersId="197638">Marques Margarida</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7023" PersId="197634">Marques Pedro</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6839" PersId="197441">Matić</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6322" PersId="124691">Mavrides</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6915" PersId="197522">Maxová</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7272" PersId="204418">Mebarek</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6890" PersId="197497">Mikser</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6933" PersId="197540">Miller</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6351" PersId="124722">Molnár</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5466" PersId="28347">Moreno Sánchez</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6433" PersId="124799">Moretti</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7310" PersId="238674">Bielowski</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6691" PersId="88882">Negrescu</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7298" PersId="233862">Nemec</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6418" PersId="124784">Nica</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6473" PersId="124836">Noichl</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7308" PersId="237779">Ohlsson</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7221" PersId="197838">Olekas</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6323" PersId="124692">Papadakis Demetris</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7324" PersId="244571">Papandreou</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7228" PersId="197845">Penkova</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6261" PersId="112744">Picula</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6982" PersId="197590">Pisapia</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5731" PersId="38595">Plumb</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6136" PersId="96998">Regner</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7296" PersId="229839">Repasi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7288" PersId="218347">Reuten</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7186" PersId="197801">Roberti</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6649" PersId="125043">Rodríguez-Piñero</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6974" PersId="197584">Rónai</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7177" PersId="197792">Rondinelli</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7269" PersId="204413">Ros Sempere</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5842" PersId="96653">Roth Neveďalová</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7326" PersId="245858">Rudner</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7128" PersId="127096">Ruiz Devesa</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6415" PersId="124781">Sant</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7039" PersId="197650">Santos</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5713" PersId="37312">Schaldemose</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6475" PersId="124837">Schuster</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7065" PersId="197677">Sidl</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6381" PersId="124747">Silva Pereira</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6093" PersId="96932">Sippel</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7176" PersId="197791">Smeriglio</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6490" PersId="124852">Stanishev</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6623" PersId="125020">Tang</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7140" PersId="197756">Tax</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5459" PersId="28340">Toia</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7247" PersId="198183">Tudose</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6336" PersId="124705">Ujhelyi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7195" PersId="197810">Ušakovs</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4990" PersId="5729">Van Brempt</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7313" PersId="239257">Variati</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7227" PersId="197844">Vitanov</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6735" PersId="185619">Wölken</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7225" PersId="197842">Yoncheva</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6373" PersId="124739">Zorrinho</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="The Left">
            <PoliticalGroup.Member.Name MepId="6926" PersId="197533">Aubry</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6604" PersId="124992">Björk</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7126" PersId="197743">Kokkalis</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6713" PersId="130833">Kouloglou</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7087" PersId="197699">Kountoura</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7279" PersId="205452">MacManus</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6354" PersId="24505">Maurel</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7303" PersId="236053">Mesure</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7190" PersId="197805">Modig</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6250" PersId="30482">Omarjee</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6920" PersId="197527">Pelletier</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7169" PersId="197785">Rodríguez Palop</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6715" PersId="131507">Urbán Crespo</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6961" PersId="197571">Villumsen</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="Verts/ALE">
            <PoliticalGroup.Member.Name MepId="7263" PersId="204335">Alametsä</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6845" PersId="197448">Andresen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5291" PersId="28161">Auken</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6905" PersId="197512">Biteau</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6837" PersId="197439">Boeselager</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6829" PersId="197431">Breyer</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6867" PersId="197470">Bricmont</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5913" PersId="96739">Bütikofer</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6964" PersId="197574">Carême</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6847" PersId="86793">Cavazzini</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6897" PersId="197503">Cormand</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6494" PersId="124856">Corrao</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7043" PersId="197654">Cuffe</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6730" PersId="183338">Dalunde</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6472" PersId="124835">D'Amato</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6924" PersId="197531">Delbos-Corfield</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6040" PersId="96868">Delli</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6850" PersId="197453">Deparnay-Grunenberg</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5899" PersId="96725">Eickhout</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6851" PersId="106936">Freund</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7294" PersId="229352">Gallée</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6853" PersId="183916">Geese</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7274" PersId="204420">Gruffat</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7034" PersId="197645">Guerreiro</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6854" PersId="197457">Hahn Henrike</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5926" PersId="96752">Häusling</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="4535" PersId="2054">Hautala</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6856" PersId="197459">Herzberger-Fofana</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6797" PersId="197398">Holmgren</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5914" PersId="96740">Jadot</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7220" PersId="197837">Jakeliūnas</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5908" PersId="96734">Keller Ska</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6939" PersId="197546">Kolaja</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6795" PersId="197395">Kuhnke</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6857" PersId="197460">Lagodinsky</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5838" PersId="96648">Lamberts</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6858" PersId="197461">Langensiepen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6859" PersId="197462">Marquardt</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7281" PersId="208722">Matthieu</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6776" PersId="193292">Metz</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6251" PersId="24942">Miranda</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6861" PersId="197464">Neumann</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7240" PersId="197889">O'Sullivan</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6863" PersId="197466">Paulus</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6482" PersId="124844">Pedicini</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6932" PersId="197539">Peksa</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6963" PersId="197573">Peter-Hansen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6529" PersId="103381">Reintke</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7217" PersId="197832">Riba i Giner</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7280" PersId="206158">Ripa</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5917" PersId="96743">Rivasi</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6900" PersId="197506">Roose</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6694" PersId="125214">Ropė</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6893" PersId="197500">Satouri</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6822" PersId="197423">Semsrott</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6738" PersId="185974">Solé</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6943" PersId="197550">Spurek</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7055" PersId="58766">Ştefănuță</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7156" PersId="197772">Strik</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6894" PersId="97236">Toussaint</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6591" PersId="124972">Urtasun</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6551" PersId="124934">Vana</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7237" PersId="197870">Van Sparrentak</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6848" PersId="197451">von Cramon-Taubadel</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6757" PersId="190464">Waitz</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7064" PersId="197675">Wiener</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
    </Result.For>
    <Result.Against Number="12">
        <Result.PoliticalGroup.List Identifier="ECR">
            <PoliticalGroup.Member.Name MepId="7278" PersId="204733">Rookmaker</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="NI">
            <PoliticalGroup.Member.Name MepId="6832" PersId="197434">Buschmann</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6630" PersId="125025">de Graaff</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7116" PersId="197734">Nikolaou-Alavanos</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6676" PersId="125093">Papadakis Kostas</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6471" PersId="124834">Sonneborn</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5602" PersId="28619">Ždanoka</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="The Left">
            <PoliticalGroup.Member.Name MepId="7113" PersId="197731">Daly</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7133" PersId="88715">Gusmão</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6628" PersId="125023">Hazekamp</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5992" PersId="96820">Matias</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7202" PersId="197817">Wallace</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
    </Result.Against>
    <Result.Abstention Number="21">
        <Result.PoliticalGroup.List Identifier="ECR">
            <PoliticalGroup.Member.Name MepId="7309" PersId="238639">Nissinen</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6808" PersId="123562">Weimers</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="ID">
            <PoliticalGroup.Member.Name MepId="6949" PersId="197556">David</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="NI">
            <PoliticalGroup.Member.Name MepId="7189" PersId="197804">Giarrusso</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7242" PersId="197935">Sinčić</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
        <Result.PoliticalGroup.List Identifier="The Left">
            <PoliticalGroup.Member.Name MepId="7089" PersId="197701">Arvanitis</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6873" PersId="187917">Botenga</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6865" PersId="197468">Demirel</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6024" PersId="96852">Ernst</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6597" PersId="124985">Flanagan</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6815" PersId="197416">Georgiou</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6814" PersId="197415">Kizilyürek</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5039" PersId="23699">Konečná</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6310" PersId="120478">Michels</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5596" PersId="28586">Papadimoulis</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7138" PersId="197754">Pereira Sandra</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6727" PersId="136236">Pimenta Lopes</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7213" PersId="197828">Pineda</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="7069" PersId="197681">Rego</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="6758" PersId="190517">Schirdewan</PoliticalGroup.Member.Name>
            <PoliticalGroup.Member.Name MepId="5837" PersId="96646">Scholz</PoliticalGroup.Member.Name>
        </Result.PoliticalGroup.List>
    </Result.Abstention>
    <Intentions Code="100">
        <Intentions.Result.For>
            <Member.Name MepId="7133" PersId="88715">Gusmão</Member.Name>
            <Member.Name MepId="5992" PersId="96820">Matias</Member.Name>
            <Member.Name MepId="5596" PersId="28586">Papadimoulis</Member.Name>
        </Intentions.Result.For>
    </Intentions>
</RollCallVote.Result>
	`
	// Create an instance of the top-level struct
	var vote RollCallVote

	// Use strings.NewReader() to create an io.Reader from the XML data
	reader := strings.NewReader(xmlData)

	// Use xml.NewDecoder() to create a new XML decoder from the reader
	decoder := xml.NewDecoder(reader)

	// Use xml.Unmarshal() to parse the XML into the struct
	if err := decoder.Decode(&vote); err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	// Now, "vote" contains the parsed XML data in Go structs
	fmt.Printf("Vote Date: %s\n", vote.Result.Date)
	fmt.Printf("Description: %s\n", vote.Result.Description.Text)
	fmt.Printf("For Political Group: %s\n", vote.Result.For.PoliticalGroupList[0].Identifier)
	// Add more fields as needed

}
