import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import debounce from "lodash.debounce";

export function resizeWidth(mainId: string, leftId: string, splitterId: string, rightId: string,
                            leftMin: number, rightMin: number, gap: number): boolean {
    const main = document.getElementById(mainId) as any;
    const left = document.getElementById(leftId) as any;
    const splitter = document.getElementById(splitterId) as any;
    const content = document.getElementById(rightId) as any;

    if (!splitter) return false

    // 鼠标按下事件
    splitter.onmousedown = function (e) {
        //色彩高亮
        splitter.classList.add('active');
        const startX = e.clientX - gap;

        // 鼠标拖动事件
        document.onmousemove = function (e) {
            splitter.left = startX;
            const endX = e.clientX - gap;

            let moveLen = splitter.left + (endX - startX); // （endx-startx）=挪动的间隔。splitter.left+挪动的间隔=右边区域最初的宽度
            const maxT = main.clientWidth - splitter.offsetWidth; // 容器宽度 - 右边区域的宽度 = 左边区域的宽度
            if (moveLen < leftMin) moveLen = leftMin; // 右边区域的最小宽度
            if (moveLen > maxT - rightMin) moveLen = maxT - rightMin; //左边区域最小宽度
            splitter.style.left = moveLen; // 设置左侧区域的宽度

            left.style.width = (moveLen / document.body.clientWidth) * 100 + '%';
            content.style.width = ((main.clientWidth - moveLen) / document.body.clientWidth - 0.008) * 100 + '%';
        };

        // 鼠标松开事件
        document.onmouseup = function (evt) {
            splitter.classList.remove('active'); //色彩复原

            document.onmousemove = null;
            document.onmouseup = null;
            splitter.releaseCapture && splitter.releaseCapture(); //当你不在须要持续取得鼠标音讯就要应该调用ReleaseCapture()开释掉
        };

        splitter.setCapture && splitter.setCapture(); //该函数在属于以后线程的指定窗口里设置鼠标捕捉
        return false;
    };

    return true
}

export function resizeHeight(contentId: string, topId: string, splitterId: string, bottomId: string,
                            topMin: number, bottomMin: number, gap: number): boolean {

    const content = document.getElementById(contentId) as any;
    const top = document.getElementById(topId) as any;
    const splitter = document.getElementById(splitterId) as any;
    const bottom = document.getElementById(bottomId) as any;

    if (!splitter) return false

    // 鼠标按下事件
    splitter.onmousedown = function (e) {
        //色彩高亮
        splitter.classList.add('active');
        const startY = e.clientY - gap;

        // 鼠标拖动事件
        document.onmousemove = function (e) {
            splitter.top = startY;
            const endY = e.clientY - gap;

            let moveLen = splitter.top + (endY - startY); // （endY-startY）=挪动的间隔。splitter.top+挪动的间隔=上边区域最初的高度
            const maxT = content.clientHeight - splitter.offsetHeight; // 容器高度 - 下边区域的宽度 = 上边区域的宽度
            if (moveLen < topMin) moveLen = topMin; // 下边区域的最小宽度
            if (moveLen > maxT - bottomMin) moveLen = maxT - bottomMin; //上边区域最小高度
            splitter.style.top = moveLen; // 设置上边区域的高度

            top.style.height = (moveLen / content.clientHeight) * 100 + '%';
            bottom.style.height = ((content.clientHeight - moveLen) / content.clientHeight - 0.008) * 100 + '%';

            resizeHandler()
        };

        // 鼠标松开事件
        document.onmouseup = function (evt) {
            splitter.classList.remove('active'); //色彩复原

            document.onmousemove = null;
            document.onmouseup = null;
            splitter.releaseCapture && splitter.releaseCapture(); //当你不在须要持续取得鼠标音讯就要应该调用ReleaseCapture()开释掉
        };

        splitter.setCapture && splitter.setCapture(); //该函数在属于以后线程的指定窗口里设置鼠标捕捉
        return false;
    };

    return true
}

export const resizeHandler = debounce(() => {
    bus.emit(settings.eventEditorContainerHeightChanged, '')
}, 50);

export function hasClass( elements, cName ){
    if (!elements) return false
    return !!elements.className.match( new RegExp( "(\\s|^)" + cName + "(\\s|$)") )
}
export function addClass(elements, cName){
    if (!elements) return
    if( !hasClass( elements,cName ) ){
        elements.className += " " + cName
    }
}
export function removeClass( elements, cName ){
    if (!elements) return
    if( hasClass( elements,cName ) ){
        elements.className = elements.className.replace( new RegExp( "(\\s|^)" + cName + "(\\s|$)" ), " " )
    }
}

export function scroll(id: string): void {
    const elem = document.getElementById(id)
    if (elem) {
        setTimeout(function(){
            elem.scrollTop = elem.scrollHeight + 100;
        },300);
    }
}

export function formatXml(xml: any) : string {
    const PADDING = ' '.repeat(2);
    const reg = /(>)(<)(\/*)/g;
    let pad = 0;

    xml = xml.replace(reg, '$1\r\n$2$3');

    return xml.split('\r\n').map((node: any, index: number) => {
        let indent = 0;
        if (node.match(/.+<\/\w[^>]*>$/)) {
            indent = 0;
        } else if (node.match(/^<\/\w/) && pad > 0) {
            pad -= 1;
        } else if (node.match(/^<\w[^>]*[^/]>.*$/)) {
            indent = 1;
        } else {
            indent = 0;
        }

        pad += indent;

        return PADDING.repeat(pad - indent) + node;
    }).join('\r\n');
}

export function getResultCls (result) {
    if (!result) return 'dp-color-unknown'

    result = result.toLowerCase()
    if (result === 'pass') {
        return 'dp-color-pass'
    } else if (result === 'fail') {
        return 'dp-color-fail'
    } else {
        return 'dp-color-unknown'
    }
}