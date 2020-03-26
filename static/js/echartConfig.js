/**
 *@Description: echart图表配置
 *@Author: imi
 *@date: 2020/3/26
 */
//柱状图
let o1 = {
    xAxis: {
        type: 'category',
        data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    yAxis: {
        type: 'value'
    },
    series: [{
        data: [120, 200, 150, 80, 70, 110, 130],
        type: 'bar',
        showBackground: true,
        backgroundStyle: {
            color: 'rgba(220, 220, 220, 0.8)'
        }
    }]
};

//仪表盘
let o2 = {
    tooltip: {
        formatter: '{a} <br/>{b} : {c}%'
    },
    series: [
        {
            name: '经验回响度',
            type: 'gauge',
            detail: {formatter: '{value}%'},
            data: [{value: 50}]
        }
    ]
};

//雷达图
let o3 = {
    tooltip: {},
    legend: {
        data: ['经验指数']
    },
    radar: {
        // shape: 'circle',
        name: {
            textStyle: {
                color: '#fff',
                backgroundColor: '#999',
                borderRadius: 3,
                padding: [1, 3]
            }
        },
        indicator: [
            {name: '优质', max: 10},
            {name: '活跃', max: 10},
            {name: '原创', max: 10},
            {name: '互动', max: 10},
            {name: '财富', max: 100000},
            {name: '现金', max: 500}
        ]
    },
    series: [{
        type: 'radar',
        // areaStyle: {normal: {}},
        name: "经验详情",
        data: [
            {
                value: [1, 1, 1, 1, 1, 1],
            },
        ]
    }]
};
let app = {}
let posList = [
    'left', 'right', 'top', 'bottom',
    'inside',
    'insideTop', 'insideLeft', 'insideRight', 'insideBottom',
    'insideTopLeft', 'insideTopRight', 'insideBottomLeft', 'insideBottomRight'
];
app.configParameters = {
    rotate: {
        min: -90,
        max: 90
    },
    align: {
        options: {
            left: 'left',
            center: 'center',
            right: 'right'
        }
    },
    verticalAlign: {
        options: {
            top: 'top',
            middle: 'middle',
            bottom: 'bottom'
        }
    },
    position: {
        options: echarts.util.reduce(posList, function (map, pos) {
            map[pos] = pos;
            return map;
        }, {})
    },
    distance: {
        min: 0,
        max: 100
    }
};

app.config = {
    rotate: 35,
    align: 'center',
    verticalAlign: 'middle',
    position: 'top',
    distance: 15,
    font:8,
    onChange: function () {
        var labelOption = {
            normal: {
                rotate: app.config.rotate,
                align: app.config.align,
                verticalAlign: app.config.verticalAlign,
                position: app.config.position,
                distance: app.config.distance
            }
        };
        myChart.setOption({
            series: [{
                label: labelOption
            }, {
                label: labelOption
            }, {
                label: labelOption
            }, {
                label: labelOption
            }]
        });
    }
};
let labelOption = {
    show: true,
    position: app.config.position,
    distance: app.config.distance,
    align: app.config.align,
    verticalAlign: app.config.verticalAlign,
    rotate: app.config.rotate,
    formatter: '{c}  {name|{a}}',
    fontSize: 16,
    rich: {
        name: {
            textBorderColor: '#fff'
        }
    }
};

// 经验统计柱状图
let o4 = {
    color: ['#674bd1', '#43d699', '#4cabce', '#e55f60'],
    tooltip: {
        trigger: 'axis',
        axisPointer: {
            type: 'shadow'
        }
    },
    legend: {
        data: ['浏览', '投票', '喜爱', '优质']
    },
    toolbox: {
        show: true,
        orient: 'vertical',
        left: 'right',
        top: 'center',
        feature: {
            mark: {show: true},
            dataView: {show: true, readOnly: false},
            magicType: {show: true, type: ['line', 'bar', 'stack', 'tiled']},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    xAxis: [
        {
            type: 'category',
            axisTick: {show: false},
            data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日',]
        }
    ],
    yAxis: [
        {
            type: 'value'
        }
    ],
    series: [
        {
            name: '浏览',
            type: 'bar',
            barGap: 0,
            label: labelOption,
            data: [320, 332, 101, 334, 390, 123, 34]
        },
        {
            name: '投票',
            type: 'bar',
            label: labelOption,
            data: [20, 12, 11, 24, 0, 34]
        },
        {
            name: '喜爱',
            type: 'bar',
            label: labelOption,
            data: [0, 32, 1, 14, 19, 6]
        },
        {
            name: '优质',
            type: 'bar',
            label: labelOption,
            data: [9, 7, 11, 9, 0, 1, 7]
        }
    ]
};
