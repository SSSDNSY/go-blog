$(function () {
    //经验页面加载完成
    let param = window.localStorage.getItem("BD_DATA")
    if(param && window.location.pathname =="/bdjy"){
        refreshData(JSON.parse(param))
    }
    $("#card2").css("display","none")
    $("#card3").css("display","none")
})

function onMyCookie(e) {
    if (e && $.trim($(e).val()).length == 0 && event.keyCode === 13) {
        alert("请输入百度cookie中的 BDUID ！")
        return
    }
    if (event.keyCode === 13 || event.type === "click") {
        myRadarChart.showLoading()
        $.post("/bdjy/api", {
            "uuid": $.trim($(e).val())
        }, function (data) {
            // { name: '优质', max: 10},
            // { name: '活跃', max: 10},
            // { name: '原创', max: 10},
            // { name: '互动', max: 10},
            // { name: '财富', max: 100000},
            // { name: '现金', max: 1000}

            // Uname: "↵谁是谁的那谁吖↵"
            // Level: "↵当前等级：经验大神 Lv1↵"
            // Intro: "3.03"
            // Expnum: "559"
            // Fans: "9"
            // Returns: "2.51"
            // Quality: "3.03"
            // Interact: "9.59"
            // Cash: "343.41"
            // Wealth: "49395"
            // Active: "9.36"
            // Origin: "10"
            // Timing: ""
            // console.log(data)
            // window["BD_DATA"] = data
            window.localStorage.setItem("BD_DATA",JSON.stringify(data))
            refreshData(data)
        });
    }
}

function showMyTotal(e){
    $.post("/bdjy/total", {
        "uuid": $.trim($(e).val())
    }, function (data) {
        console.log(data)
    })
}

function refreshData(data) {
    $(avatorId).attr("src", data.Others.avatorUrl)
    $("#bdname").text(data.Uname)
    $("#bdLevel").text(data.Level.substr(6))
    $("#bdLevelDes").text("（篇数"+data.Expnum+" ，粉丝"+data.Fans+"）")
    o3.series[0].data[0] = [data.Quality, data.Active, data.Origin, data.Interact, data.Wealth, data.Cash]
    myRadarChart.setOption(o3)
    // { name: '现金', max: 1000}
    myRadarChart.hideLoading()
}

function changeCard(p) {
    $(p).addClass("active").addClass("text-info").parent().siblings().children().removeClass("active").removeClass("text-info");
    $("#card"+p.dataset["card"]).show().siblings("div").hide()
}