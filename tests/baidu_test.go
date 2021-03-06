package test

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"go-blog/util"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGetPerson(t *testing.T) {
	t1 := time.Now()
	html, err := util.GetPerson("BAIDUID=CF509AA1B0B95282EAFB0D768A96EBBF:FG=1; BIDUPSID=CF509AA1B0B95282EAFB0D768A96EBBF; PSTM=1563796512; bdshare_firstime=1563798260040; 9f63fb91e9b26388400f0e81=1; 7f766dafaaabe04101e1d0a3=1; 636f38bbe8045c96b84610f4=1; 4d58d5415c021c9dd4e9c0e0=1; isStepGuided=done; cd4c297923ae37756e6e60d3=1; EXP_TOPMSG0=2; 8cdccae9023be3315513cd42=1; 22fe7ced1e2d6d3002617ff8=1; f0062228d23aa8fbd3f0c8fc=1; 75ab0bcb8987e2d6864db2b8=1; b2c186c821b055c46ff6ff68=1; 15622f24704b10fdfdbea56f=1; 92255446fe9cbac51748f472=1; 148a1921e19e0a4d71c3b1da=1; 9225544670cbc8851748f469=1; 2fb0ba4005b48840f3ec5f6a=1; afd8f4de3e177774e286e9ac=1; fea4511ac83692f7bb912521=1; 47a29f2461c340c0142399ae=1; 4ae03de3f603e53eff9e6ba7=1; magazine18153=1; 380abd0aba29c35d91192c36=1; e75057f21996d5ebc91a890d=1; 6d704a1331b05228da51ca69=1; a24b33cda0956759ff002b5a=1; ae97a646005904bbfd461d90=1; d621e8dae4173b2865913fb5=1; 3d69c55111da40f0cf02d733=1; magazine18000=1; ab0b56309751bac15afa7d21=1; 4d58d54168c3089dd4e9c0a1=1; 48a42057f621ada924250433=1; 3c48dd34c5accbe10be35807=1; magazine18610=1; magazine7979=1; ac6a9a5e29cc842b643eac6d=1; 7908e85c44a98fef491ad270=1; 00a07f3858c32382d128dc5f=1; f71d6037bf0e231ab641d1b3=1; BDSFRCVID=AL8OJeCmH6lrQNjuXdn3rnU4YmKK0gOTHllvxDjaaiNdIqIVJeC6EG0Ptf8g0KubTYO3ogKK0gOTH6KF_2uxOjjg8UtVJeC6EG0Ptf8g0M5; H_BDCLCKID_SF=tJPq_C05tDP3eb5m-trSqRIeh2T22-uXKK_sWpccBhcqEn6Sj4JW2RIg5PbMhtoJQKbr0J6cWKJJ8UbShb7OKt0FBnoeJhIJbbnpaJ5nJq5nhMJmb67JDMP0-x5Rthcy523ion3vQpP-OpQ3DRoWXPIqbN7P-p5Z5mAqKl0MLPbtbb0xXj_0D63BjHtfJjks5ITJ3RTsHtcoD5rnhPF_-P6MW4brWMT-0Kj7_j6l0l_bfPPl0-7hLlRbyxciB5OMBan7WnRI3Kn6otbJX5j-3q4ZyxomtfQxtNRJ0DnjtpvhHRoejprobUPUDMJ9LUkqW2cdot5yBbc8eIna5hjkbfJBQttjQn3hfIkj2CKLtD0aMK_ljT03KP4E52Ty2t70aDTbW-38HJrqfKvdKMRcy4LdhtAqLJbyLeQfWPc-MP5I_In_hfRvD--g3-7Q2n5QtT6-oIokyRnIsxQ_bf--QfbQ0hOhqP-jW5TaKfbxLJ7JOpkxbUnxy5KUQRPH-Rv92DQMVU52QqcqEIQHQT3m5-5bbN3ut6IHJbCJoDD5tIv5DnQmq4n_h-FHbfOy--CXaDTJVhcDtp7keqOzy5j4KP_b-lJ20MjBbmrdVpcEWhk2ep72y-vSMPLuXPbAybjjJDJR-qcH0KQpsIJM5-DWbT8U5f5lLpONaKviaKJjBMb1MlvDBT5h2M4qMxtOLR3pWDTm_q5TtUJMeCnTDMFhe6jWeHDOJ5-Jf5vfL5rEMRbofnuk-PnVeU_Te-nZKxJLWncQQfjk2t5iebuRK4TljhFDyP4jKMRnWIJQLtng2JjHjqrgyJQz3xI8LNj405OTbIFO0KJcbRoxoqARhPJvyT8sXnO7tfnlXbrtXp7_2J0WStbKy4oTjxL1Db3JKjvMtgDtVJO-KKCMhCIlDU5; 4d58d54144a6149dd5e9c05c=1; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; 22fe7cedf5b0323003617f6b=1; 456c463bba74164b583144e9=1; 4853e1e548a25b1909f726a9=1; b2c186c803ba4bc46ff6ff6e=1; ab69b2706e231f2ca7189fb5=1; ad310e809a275b5849f49ede=1; e4511cf3ef60ae6a855eaf3e=1; 574c5219e7c8302c8c9dc15b=1; BDUSS=UhiNDJWQlhpMGhpbFlxOUY0UUlhdmJsTmgyUVpwS0tzN1BxOFF3UWtBSDd3VzFlSUFBQUFBJCQAAAAAAAAAAAEAAABUZ7REy63Kx8uttcTEx8ut37kAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPs0Rl77NEZeO; __cfduid=d37d2fbba930826a7c3b6c5fedbb323c01582125790; aa6a2c140da8574c4c19c4a4=1; 870c6fc3dad5a0b03ee4be5a=1; a65957f4af379524e77f9b46=1; e5c39bf56a8db639d6603379=1; fc07f989bc32c412ffe51900=1; 20095761792c128a0621b418=1; 1612d5001fc8c7a30f1eee39=1; 6fb756ecc0f54f651858fbed=1; b24f6c82e671f7c7bfe5da8b=1; 4ae03de342c9b17fff9e6bae=1; 2d5afd69b78ef7c4a3e28e78=1; SHOW_VIDEO_GUIDE_2020=1; 93f9803fd78121a0e46f5593=1; MCITY=-%3A; d5a880ebc8a5dc13f147cc86=1; 54b6b9c036564d6c593b475a=1; cb5d6105578fc3415d2fe00e=1; e75aca858534c6542edac6f1=1; f79b7cb32d7cf09144023ec5=1; BD_BOXFO=_uvqfguNvoGWC; 9c69d48fe813c813c9024e05=1; PS_REFER=0; Hm_lvt_46c8852ae89f7d9526f0082fafa15edd=1584578594,1584607650,1584608171,1584611302; H_PS_PSSID=30975_1455_31124_21100_30908_30824_31086_26350_22160; Hm_lpvt_46c8852ae89f7d9526f0082fafa15edd=1584618513")
	fmt.Println(html, err)
	t2 := time.Now()
	fmt.Println("请求百度耗时=", t2.Sub(t1).String())
}

func TestParsePerson(t *testing.T) {
	html, _ := util.GetPerson("BAIDUID=CF509AA1B0B95282EAFB0D768A96EBBF:FG=1; BIDUPSID=CF509AA1B0B95282EAFB0D768A96EBBF; PSTM=1563796512; bdshare_firstime=1563798260040; 9f63fb91e9b26388400f0e81=1; 7f766dafaaabe04101e1d0a3=1; 636f38bbe8045c96b84610f4=1; 4d58d5415c021c9dd4e9c0e0=1; isStepGuided=done; cd4c297923ae37756e6e60d3=1; EXP_TOPMSG0=2; 8cdccae9023be3315513cd42=1; 22fe7ced1e2d6d3002617ff8=1; f0062228d23aa8fbd3f0c8fc=1; 75ab0bcb8987e2d6864db2b8=1; b2c186c821b055c46ff6ff68=1; 15622f24704b10fdfdbea56f=1; 92255446fe9cbac51748f472=1; 148a1921e19e0a4d71c3b1da=1; 9225544670cbc8851748f469=1; 2fb0ba4005b48840f3ec5f6a=1; afd8f4de3e177774e286e9ac=1; fea4511ac83692f7bb912521=1; 47a29f2461c340c0142399ae=1; 4ae03de3f603e53eff9e6ba7=1; magazine18153=1; 380abd0aba29c35d91192c36=1; e75057f21996d5ebc91a890d=1; 6d704a1331b05228da51ca69=1; a24b33cda0956759ff002b5a=1; ae97a646005904bbfd461d90=1; d621e8dae4173b2865913fb5=1; 3d69c55111da40f0cf02d733=1; magazine18000=1; ab0b56309751bac15afa7d21=1; 4d58d54168c3089dd4e9c0a1=1; 48a42057f621ada924250433=1; 3c48dd34c5accbe10be35807=1; magazine18610=1; magazine7979=1; ac6a9a5e29cc842b643eac6d=1; 7908e85c44a98fef491ad270=1; 00a07f3858c32382d128dc5f=1; f71d6037bf0e231ab641d1b3=1; BDSFRCVID=AL8OJeCmH6lrQNjuXdn3rnU4YmKK0gOTHllvxDjaaiNdIqIVJeC6EG0Ptf8g0KubTYO3ogKK0gOTH6KF_2uxOjjg8UtVJeC6EG0Ptf8g0M5; H_BDCLCKID_SF=tJPq_C05tDP3eb5m-trSqRIeh2T22-uXKK_sWpccBhcqEn6Sj4JW2RIg5PbMhtoJQKbr0J6cWKJJ8UbShb7OKt0FBnoeJhIJbbnpaJ5nJq5nhMJmb67JDMP0-x5Rthcy523ion3vQpP-OpQ3DRoWXPIqbN7P-p5Z5mAqKl0MLPbtbb0xXj_0D63BjHtfJjks5ITJ3RTsHtcoD5rnhPF_-P6MW4brWMT-0Kj7_j6l0l_bfPPl0-7hLlRbyxciB5OMBan7WnRI3Kn6otbJX5j-3q4ZyxomtfQxtNRJ0DnjtpvhHRoejprobUPUDMJ9LUkqW2cdot5yBbc8eIna5hjkbfJBQttjQn3hfIkj2CKLtD0aMK_ljT03KP4E52Ty2t70aDTbW-38HJrqfKvdKMRcy4LdhtAqLJbyLeQfWPc-MP5I_In_hfRvD--g3-7Q2n5QtT6-oIokyRnIsxQ_bf--QfbQ0hOhqP-jW5TaKfbxLJ7JOpkxbUnxy5KUQRPH-Rv92DQMVU52QqcqEIQHQT3m5-5bbN3ut6IHJbCJoDD5tIv5DnQmq4n_h-FHbfOy--CXaDTJVhcDtp7keqOzy5j4KP_b-lJ20MjBbmrdVpcEWhk2ep72y-vSMPLuXPbAybjjJDJR-qcH0KQpsIJM5-DWbT8U5f5lLpONaKviaKJjBMb1MlvDBT5h2M4qMxtOLR3pWDTm_q5TtUJMeCnTDMFhe6jWeHDOJ5-Jf5vfL5rEMRbofnuk-PnVeU_Te-nZKxJLWncQQfjk2t5iebuRK4TljhFDyP4jKMRnWIJQLtng2JjHjqrgyJQz3xI8LNj405OTbIFO0KJcbRoxoqARhPJvyT8sXnO7tfnlXbrtXp7_2J0WStbKy4oTjxL1Db3JKjvMtgDtVJO-KKCMhCIlDU5; 4d58d54144a6149dd5e9c05c=1; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; 22fe7cedf5b0323003617f6b=1; 456c463bba74164b583144e9=1; 4853e1e548a25b1909f726a9=1; b2c186c803ba4bc46ff6ff6e=1; ab69b2706e231f2ca7189fb5=1; ad310e809a275b5849f49ede=1; e4511cf3ef60ae6a855eaf3e=1; 574c5219e7c8302c8c9dc15b=1; BDUSS=UhiNDJWQlhpMGhpbFlxOUY0UUlhdmJsTmgyUVpwS0tzN1BxOFF3UWtBSDd3VzFlSUFBQUFBJCQAAAAAAAAAAAEAAABUZ7REy63Kx8uttcTEx8ut37kAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPs0Rl77NEZeO; __cfduid=d37d2fbba930826a7c3b6c5fedbb323c01582125790; aa6a2c140da8574c4c19c4a4=1; 870c6fc3dad5a0b03ee4be5a=1; a65957f4af379524e77f9b46=1; e5c39bf56a8db639d6603379=1; fc07f989bc32c412ffe51900=1; 20095761792c128a0621b418=1; 1612d5001fc8c7a30f1eee39=1; 6fb756ecc0f54f651858fbed=1; b24f6c82e671f7c7bfe5da8b=1; 4ae03de342c9b17fff9e6bae=1; 2d5afd69b78ef7c4a3e28e78=1; SHOW_VIDEO_GUIDE_2020=1; 93f9803fd78121a0e46f5593=1; MCITY=-%3A; d5a880ebc8a5dc13f147cc86=1; 54b6b9c036564d6c593b475a=1; cb5d6105578fc3415d2fe00e=1; e75aca858534c6542edac6f1=1; f79b7cb32d7cf09144023ec5=1; BD_BOXFO=_uvqfguNvoGWC; 9c69d48fe813c813c9024e05=1; PS_REFER=0; Hm_lvt_46c8852ae89f7d9526f0082fafa15edd=1584578594,1584607650,1584608171,1584611302; H_PS_PSSID=30975_1455_31124_21100_30908_30824_31086_26350_22160; Hm_lpvt_46c8852ae89f7d9526f0082fafa15edd=1584618513")

	Bd, _ := util.ParsePerson(html)

	fmt.Printf("%+v\n", Bd)

}

//已发布经验测试
func TestGetExpNum(t *testing.T) {
	util.ParseExPublished(BDID)
}

//已悬赏经验测试
func TestRewardExpNum(t *testing.T) {
	util.ParRewardExp()
}

//log
func TestLogFetal(t *testing.T) {
	logs.Error("123")
	logs.Error("qqq")
	logs.Error("4444444444")
	logs.Error("34444")
}

func TestGoQuery(t *testing.T) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(TEST_HTML_TXT))
	var view = 0
	var favo = 0
	var vote = 0

	doc.Find("p.f-meta").Each(func(i int, s *goquery.Selection) {
		_v, _ := strconv.Atoi(s.Find("span.view-count").Text())
		_t, _ := strconv.Atoi(s.Find("span.vote-count").Text())
		_f, _ := strconv.Atoi(s.Find("span.favo-count").Text())
		view += _v
		vote += _t
		favo += _f
	})
	rtnMap := make(map[string]string)
	rtnMap["v"] = strconv.Itoa(view)
	rtnMap["t"] = strconv.Itoa(vote)
	rtnMap["f"] = strconv.Itoa(favo)

	fmt.Printf("%v", rtnMap)
}
