import http from 'k6/http';
import { check } from 'k6';

export let options = {
  stages: [
    { duration: '10s', target: 10 },
    { duration: '20s', target: 30 },
    { duration: '10s', target: 0 },
  ],
};
export default function () {
  var url = 'TOREPLACE';
  var payload = `
第二次世界大戰鑾披汶·頌堪的法西斯政府傾向大日本帝國，1941年12月7日，日本發動太平洋戰爭，日本和暹羅簽訂《日泰攻守同盟条約》，泰國加入軸心國，亦成為日本唯一在亞洲的盟友。1942年1月25日泰國宣佈向英美宣戰，日本曾將部份在緬甸和馬來亞半島北部佔領地割讓給暹羅。1945年8月15日日本投降，暹羅隨即在翌日宣佈「暹羅1942年1月25日對英美宣戰宣言無效」，暹羅的「宣戰無效」宣言被同盟國承認。1949年改名泰國。第二次世界大戰後泰國成為美國在東南亞的主要軍事盟國。在東南亞地區，泰國亦是一個舉足輕重的國家；首都曼谷是該區域中國際化程度很高的大都會區。另外，泰國是東協始創國之一，亦在東南亞區內事務有積極的參與。

二戰後的泰國長期實行軍政府獨裁統治，鑾披汶·頌堪、沙立·他那叻、他儂·吉滴卡宗、江萨·差瑪南及炳·廷素拉暖等先後掌權。1991年泰國軍事政變，軍方推翻差猜·春哈旺、憲法和國会。其後泰國人民上街遊行，抗議軍方統治。軍方血腥鎮壓，迫使泰王介入。總理蘇欽達和示威人士領袖雙雙跪在泰王面前，承諾平息風波，示威始告结束，選舉和憲法也得到恢复，泰國開始民主化。泰國民主黨的川·立派成為總理。1998年7月14日，他信·西那瓦創建泰愛泰黨，並任黨主席。不到五年期間，泰愛泰黨合併幾個小黨，形成一黨獨大的局面。

早在清朝時，華人就已經來到泰國謀生，並漸漸融入當地。泰國歷史上建立吞武里王國的就是具有潮汕血統鄭昭。如今，鄭王像仍被泰國民眾虔誠地供奉。泰國華商主要來自於廣東潮汕地區。泰國的重要經濟支柱都由華人把持，不少大型企業都是由華商創辦。華人的經濟地位相比當地泰族原住民較為優越。但因為泰國是個佛教國家，住民淳樸，對華人沒有太多的種族對立。

上世紀5、60年代，中國大陸政權更替後，中國開始向東南亞各國輸出革命[29]，支持所在國革命組織顛覆政府。因為東南亞共產革命的參加者幾乎都是華人，導致當地各國政府為防共而排華。泰國是東南亞國家中排華最輕微的，沒有像印尼那樣大屠殺。但為切斷華人與母國的血脈聯繫，泰國政府禁止教授華文和使用華文。華人為求自保，都放棄原來的中國華語姓氏，而改用有相同意思的泰語文字來當作他們的新姓氏甚或採用當地姓氏。這運動後來稱之為「改姓名運動」。

鄧小平時代，中國改善與東南亞國家的關係，在1989年，中國徹底結束了對東南亞的革命輸出，緬共與泰共失去了物質支持而解散。中國與泰國的經濟合作日益增多，泰國允許恢復華文教學。泰國經濟的起飛，也吸引了大批華人到泰國發展。華人在今天的泰國經濟和政治中，更扮演着重要的角色。泰國前總理他信、阿披實、英拉、泰國最大的商業集團正大集團的董事長等等都是華裔，華商在泰國的經濟發展中擁有舉足輕重的地位。

sex
`;

  let res = http.post(url, payload);
  // console.log(JSON.stringify(res, null, 4));
  check(res, { 'status was 400': r => r.status == 400 });
}
