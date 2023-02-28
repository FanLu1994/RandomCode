/**
 * @Description: 插件后台运行的脚本
 * @author: fanlu
 * @date:  2023/2/21
 * @project: chromeUserTime
 */


function init() {
    let count = 0
    // TODO:根据根域名来判断
    let currentDomain = ""
    setInterval(()=>{
        chrome.tabs.getSelected(null, function (tab) {   // 获取当前打开的页面
            if(!tab.url.startsWith("http")){
                count = 0
                chrome.browserAction.setBadgeText({
                    text:'',
                })
            }else{
                let url = tab.pendingUrl?tab.pendingUrl:tab.url
                let domain = url.split("/")
                if(domain[2]){
                    currentDomain = domain[2]
                    chrome.browserAction.setBadgeText({
                        text:count+'',
                    })
                    count++
                }
            }
        });
    },1000)
    console.log("开启监听")
    chrome.tabs.onActivated.addListener(activeInfo=>{
        chrome.tabs.get(activeInfo.tabId, tab => {
            console.log(tab)
            let url = tab.pendingUrl?tab.pendingUrl:tab.url
            let domain = url.split("/")
            console.log(domain[2])
            if(domain[2]!==currentDomain){
                count=0
            }
        })
    })
}

init()
