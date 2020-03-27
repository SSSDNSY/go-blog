$(function () {
    //经验页面加载完成
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
    let statics = JSON.parse(window.localStorage.getItem("BD_TOTAL_DATA"))
    let size = 10
    var cp = $("#currentPage")
    var ul = $("#postExps")
    switch (w) {
        case 'pre':
            let idx = Number(cp.val()) - 1;
            if (idx < 1) break;
            $("#postExps").empty()
            let ff1 = 0;
            $("#currentPage").val(idx)
            for (let [ii, expMap] of Object.entries(statics)) {
                for (let [jj, tm] of Object.entries(expMap)) {
                    if (tm.hasOwnProperty("et") && ff1++ <size * idx && ff1 >size * (idx - 1)) {
                        ul.append(generateATag('#', ff1 % 2 == 0 ? 'primary' : 'light', tm["et"]))
                    }
                }
            }
            cp.val(idx)
            break;
        case 'next':
            let idy = Number(cp.val()) + 1;
            if (Object.entries(statics).length-1 < idy) break
            $("#postExps").empty()
            let ff2 = 0;
            for (let [ii, expMap] of Object.entries(statics)) {
                for (let [jj, tm] of Object.entries(expMap)) {
                    if (tm.hasOwnProperty("et") && ff2++ < size * idy && ff2 > size * (idy - 1)) {
                        ul.append(generateATag('#', ff2 % 2 == 0 ? 'primary' : 'light', tm["et"]))
                    }
                }
            }
            cp.val(idy)
            break;
        case 'goto':
            if (event.keyCode === 13) {
                // if ( typeof cp.val() != "number")break; TODO 判断字符串是否是数字
                let idz = Number(cp.val());
                if(idz<1 || idz> Object.entries(statics).length-1 )break
                let ff3 = 0;
                $("#postExps").empty()
                for (let [ii, expMap] of Object.entries(statics)) {
                    for (let [jj, tm] of Object.entries(expMap)) {
                        if (tm.hasOwnProperty("et") && ff3++ < size*idz && ff3>size * (idz - 1)) {
                            ul.append(generateATag('#', ff3 % 2 == 0 ? 'primary' : 'light', tm["et"]))
                        }
                    }
                }
            }
            break;
        default:
            let ffi = 1;
            $("#currentPage").val("1")
            for (let [ii, expMap] of Object.entries(statics)) {
                for (let [jj, tm] of Object.entries(expMap)) {
                    if (tm.hasOwnProperty("et") && ffi++ <= size) {
                        ul.append(generateATag('#', ffi % 2 == 0 ? 'primary' : 'light', tm["et"]))
                    }
                }
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
    }
}


function changeCard(p) {
    $(p).addClass("active").addClass("text-info").parent().siblings().children().removeClass("active").removeClass("text-info");
    $("#card" + p.dataset["card"]).show().siblings("div.card").hide()
    if (p.dataset["card"] == 2 || p.dataset["card"] == 3) {
        $("#post").hide()
    } else {
        $("#post").show()
    }
}

function searchReward(param) {
    console.log(param)
}

function onMyPerExp() {

}