/**
 *@Description:
 *@Author: imi
 *@date: 2020/4/10
 */
function getWord($) {
    $.get("/about/getWord", function (data) {
        $("#pageHeader").text(data)
    });
}

getWord($)
setInterval(function (window) {
    getWord($)
}, 1000 * 15)

const PAGE_SIZE = 10


$(function () {
    //经验页面加载完成
    window.localStorage.setItem('CURRENT_PAGE', '1')
    let param = window.localStorage.getItem("BD_PERSON_DATA")
    let statics = window.localStorage.getItem("BD_TOTAL_DATA")
    if (param && window.location.pathname == "/bdjy") {
        refreshData(JSON.parse(param))
        refreshStatics(JSON.parse(statics))
        doPage()
    }
    $("#card2").css("display", "none")
    $("#card3").css("display", "none")
})

function onMyCookie(e) {
    if (e && $.trim($(e).val()).length == 0 && event.keyCode === 13) {
        alert("请输入百度cookie中的 BDUID ！")
        return
    }
    if (event.keyCode === 13 || event.type === "click") {
        myStatistics.showLoading()
        myRadarChart.showLoading()
        $.post("/bdjy/person", {
            "uuid": $.trim($(e).val())
        }, function (data) {
            window.localStorage.setItem("BD_PERSON_DATA", JSON.stringify(data))
            refreshData(data)
        });
    }
    showStatic(e)
}

function showStatic(e) {
    $.post("/bdjy/static", {
        "uuid": $.trim($(e).val())
    }, function (data) {
        console.log(data)
        window.localStorage.setItem("BD_TOTAL_DATA", JSON.stringify(data))
        myStatistics.hideLoading()
        refreshStatics(data)
    })
}

function refreshData(data) {
    $(avatorId).attr("src", data.Others.avatorUrl)
    $("#bdname").text(data.Uname)
    $("#bdLevel").text(data.Level.substr(6))
    $("#bdLevelDes").text("（篇数" + data.Expnum + " ，粉丝" + data.Fans + "）")
    o3.series[0].data[0] = [data.Quality, data.Active, data.Origin, data.Interact, data.Wealth, data.Cash]
    myRadarChart.setOption(o3)
    // { name: '现金', max: 1000}
    myRadarChart.hideLoading()
}

function refreshStatics(data) {

    //计算展示数据
    let ds = caleStaticData(data)
    //刷新统计条形图
    refreshShart(ds)
    //经验排行

}

function doPage(w, t) {

    var curPage = window.localStorage.getItem('CURRENT_PAGE');
    var cp = $("#currentPage")
    var postExps = $("#postExps")
    var rewardList = $("#rewardList")

    switch (curPage) {
        case '1':
            let statics = JSON.parse(window.localStorage.getItem("BD_TOTAL_DATA"))
            $("#totalPage").val("/   "+((2 * Object.entries(statics).length)-1));

            switch (w) {
                case 'pre':
                    let idx = Number(cp.val()) - 1;
                    if (idx < 1) break;
                    $("#postExps").empty()
                    let ff1 = 0;
                    for (let [ii, expMap] of Object.entries(statics)) {
                        for (let [jj, tm] of Object.entries(expMap)) {
                            if (tm.hasOwnProperty("et") && (PAGE_SIZE * (idx - 1)) <= ff1++ && ff1 <= (idx * PAGE_SIZE)) {
                                postExps.append(generateATag('#', ff1 % 2 == 0 ? 'primary' : 'light', tm["et"]))
                            }
                        }
                    }
                    cp.val(idx)
                    break;
                case 'next':
                    let idy = Number(cp.val()) + 1;
                    if (2 * Object.entries(statics).length -1 < idy) break
                    $("#postExps").empty()
                    let ff2 = 0;
                    for (let [ii, expMap] of Object.entries(statics)) {
                        for (let [jj, tm] of Object.entries(expMap)) {
                            if (tm.hasOwnProperty("et") && (PAGE_SIZE * (idy - 1)) <= ff2++ && ff2 <= (idy * PAGE_SIZE)) {
                                postExps.append(generateATag('#', ff2 % 2 == 0 ? 'primary' : 'light', tm["et"]))
                            }
                        }
                    }
                    cp.val(idy)
                    break;
                case 'goto':
                    if (event.keyCode === 13) {
                        // if ( typeof cp.val() != "number")break; TODO 判断字符串是否是数字
                        let idz = Number(cp.val());
                        if (idz < 1 || idz > 2 * Object.entries(statics).length-1) break
                        let ff3 = 0;
                        $("#postExps").empty()
                        for (let [ii, expMap] of Object.entries(statics)) {
                            for (let [jj, tm] of Object.entries(expMap)) {
                                if (tm.hasOwnProperty("et") && ff3++ < PAGE_SIZE * idz && ff3 >= PAGE_SIZE * (idz - 1)) {
                                    postExps.append(generateATag('#', ff3 % 2 == 0 ? 'primary' : 'light', tm["et"]))
                                }
                            }
                        }
                    }
                    break;
                default:
                    let ffi = 0;
                    $("#currentPage").val("1")
                    for (let [ii, expMap] of Object.entries(statics)) {
                        for (let [jj, tm] of Object.entries(expMap)) {
                            if (tm.hasOwnProperty("et") && ffi++ <= PAGE_SIZE) {
                                postExps.append(generateATag('#', ffi % 2 == 0 ? 'primary' : 'light', tm["et"]))
                            }
                        }
                    }
                    break;
            }
            break;
        case '2':
            let reward = JSON.parse(window.localStorage.getItem("BD_REWARD_DATA"))
            $("#card2  #totalPage").val("/   "+(Math.floor(Object.entries(reward).length / PAGE_SIZE)-1))
            switch (w) {
                case 'pre':
                    let idx = Number($("#card2  #currentPage").val()) - 1;
                    if (idx < 1) break;
                    rewardList.empty();
                    for (let p = (idx - 1) * PAGE_SIZE; (idx - 1) * PAGE_SIZE <= p && p < PAGE_SIZE * idx; p++) {
                        rewardList.append(generateATag('#', p % 2 == 0 ? 'warning' : 'dark', reward[p]))
                    }
                    $("#card2  #currentPage").val(idx)
                    break;
                case 'next':
                    let idy = Number($("#card2  #currentPage").val()) + 1;
                    if (idy >Math.floor( Object.entries(reward).length / PAGE_SIZE)-1) break;
                    rewardList.empty();
                    for (let p = idy * PAGE_SIZE ; idy * PAGE_SIZE <= p && p< (idy + 1) * PAGE_SIZE; p++) {
                        rewardList.append(generateATag('#', p % 2 == 0 ? 'warning' : 'dark', reward[p]))
                    }
                    $("#card2  #currentPage").val(idy)
                    break;
                case 'goto':
                    if (event.keyCode === 13) {
                        let idg = Number($("#card2  #currentPage").val())
                        if (idg < 1 || idg > Math.floor(Object.entries(reward).length / PAGE_SIZE)-1) break;
                        rewardList.empty();
                        for (let p = idg * PAGE_SIZE; idg * PAGE_SIZE <= p && p < (idg + 1) * PAGE_SIZE ; p++) {
                            rewardList.append(generateATag('#', p % 2 == 0 ? 'warning' : 'dark', reward[p]))
                        }
                        $("#card2  #currentPage").val(idg)
                    }
                    break;
                default:
                    break;
            }
            break;
    }


}

function generateATag(a, b, c) {
    return `<a href="${a}" class="list-group-item list-group-item-${b}">${c}</a>`
}

function caleStaticData(data) {
    let totalView = 0;
    let totalVote = 0;
    let totalFavo = 0;
    let totalHigh = 0;
    if (data && Object.entries(data).length > 0) {
        for (let [ii, expMap] of Object.entries(data)) {
            for (let [jj, tm] of Object.entries(expMap)) {
                if (tm.hasOwnProperty("eview")) {
                    totalView += Number(tm["eview"])
                    totalVote += Number(tm["evote"])
                    totalFavo += Number(tm["vfavo"].split("：")[1].split("|")[0])
                }
                if (tm.hasOwnProperty("ehq")) {
                    totalHigh++;
                }
            }
        }
    }
    return [totalView, totalVote, totalFavo, totalHigh]
}

function refreshShart(data) {

    if (data && o4 && data.length > 0) {
        o4.series.forEach(function (val, idx, arr) {
            // console.log(val + " " + idx + " " + arr[idx].data);
            arr[idx].data = [data[idx]]
        })
        myStatistics.setOption(o4)
        myStatistics.hideLoading()
    }
}


function changeCard(p) {
    $(p).addClass("active").addClass("text-info").parent().siblings().children().removeClass("active").removeClass("text-info");
    $("#card" + p.dataset["card"]).show().siblings("div.card").hide()
    if (p.dataset["card"] == 2) {
        $("#post").hide()
        showReward();
        window.localStorage.setItem('CURRENT_PAGE', '2')

    } else if (p.dataset["card"] == 3) {
        $("#post").hide()
        window.localStorage.setItem('CURRENT_PAGE', '3')

    } else {
        window.localStorage.setItem('CURRENT_PAGE', '1')
        $("#post").show()
    }
}

function showReward() {

    let rewardData = JSON.parse(window.localStorage.getItem('BD_REWARD_DATA'));
    if (rewardData && rewardData[0]&& rewardData[0].length > 0) {
        $("#rewardList").empty()
        for (let i = 0; i < PAGE_SIZE; i++) {
            $("#rewardList").append(generateATag('#', i % 2 == 0 ? 'warning' : 'dark', rewardData[i]))
        }
    } else {
        $("#rewardList").append(generateATag('#', 'warning', "悬赏数据获取中..."))
        $.post("/bdjy/reward", {}, function (data) {
            $("#rewardList").empty()
            window.localStorage.setItem("BD_REWARD_DATA", JSON.stringify(data))
            for (let i = 0; i < PAGE_SIZE; i++) {
                $("#rewardList").append(generateATag('#', i % 2 == 0 ? 'warning' : 'dark', data[i]))
            }
        });
    }

}

function doRewardPage() {

}


function searchReward(param) {
    console.log(param)
}

function onMyPerExp() {

}


