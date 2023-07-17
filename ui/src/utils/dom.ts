import bus from "@/utils/eventBus";
import settings from "@/config/settings";
import debounce from "lodash.debounce";

export function resizeWidth(mainId: string, leftId: string, splitterId: string, rightId: string,
                            leftMin: number, rightMin: number): boolean {

    const main = document.getElementById(mainId) as any;
    const left = document.getElementById(leftId) as any;
    const splitter = document.getElementById(splitterId) as any;
    const content = document.getElementById(rightId) as any;

    if (!splitter) return false

    // 鼠标按下事件
    splitter.onmousedown = function (e) {
        splitter.classList.add('active');

        const leftGap = main.getBoundingClientRect().left

        document.onmousemove = function (e) {
            let leftNewWidth = e.clientX - leftGap;

            const availableWidth = main.clientWidth - splitter.offsetWidth;

            if (leftNewWidth < leftMin) leftNewWidth = leftMin;
            if (leftNewWidth > availableWidth - rightMin) leftNewWidth = availableWidth - rightMin;

            left.style.width = leftNewWidth + 'px';
            content.style.width = availableWidth - leftNewWidth + 'px';
        };

        document.onmouseup = function (e) {
            splitter.classList.remove('active');

            document.onmousemove = null;
            document.onmouseup = null;
            splitter.releaseCapture && splitter.releaseCapture();
        };

        splitter.setCapture && splitter.setCapture();

        return false;
    };

    return true
}

export function resizeHeight(mainId: string, topId: string, splitterId: string, bottomId: string,
                            topMin: number, bottomMin: number): boolean {

    const main = document.getElementById(mainId) as any;
    const top = document.getElementById(topId) as any;
    const splitter = document.getElementById(splitterId) as any;
    const bottom = document.getElementById(bottomId) as any;

    if (!splitter) return false

    splitter.onmousedown = function (e) {
        splitter.classList.add('active');

        const topGap = main.getBoundingClientRect().top

        document.onmousemove = function (e) {
            let topNewHeight = e.clientY - topGap;

            const availableHeight = main.clientHeight - splitter.offsetHeight;

            if (topNewHeight < topMin) topNewHeight = topMin;
            if (topNewHeight > availableHeight - bottomMin) topNewHeight = availableHeight - bottomMin;

            top.style.height = topNewHeight + 'px';
            bottom.style.height = availableHeight - topNewHeight + 'px';

            resizeHandler()
        };

        document.onmouseup = function(e) {
            splitter.classList.remove('active');

            document.onmousemove = null;
            document.onmouseup = null;
            splitter.releaseCapture && splitter.releaseCapture();
        };

        splitter.setCapture && splitter.setCapture();

        return false;
    };

    return true
}

export const resizeHandler = debounce(() => {
    bus.emit(settings.eventEditorAction, {act: settings.eventTypeContainerHeightChanged})
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

export function getContextMenuStyle (x, y, height) {
    let top = y + 6
    if (y + height > document.body.clientHeight)
        top = document.body.clientHeight - height

    return {
        zIndex: 99,
        position: 'fixed',
        maxHeight: 40,
        textAlign: 'center',
        left: `${x + 10}px`,
        top: `${top}px`,
    }
}

export function getContextMenuStyle2(e) {
    console.log('getContextMenuStyle2', e.clientY)

    const style = {
        left: e.clientX + 'px',
        top: (e.clientY - 12 > 6 ? e.clientY - 12 : 6)  + 'px',
    }

    return style
}

export function getRightTabPanelPosition(tabId) {
    console.log('getRightTabPanelPosition', tabId)

    let ret = {}

    const elem = document.getElementById(tabId)

    if (elem) {
        const pos = elem.getBoundingClientRect()
        ret = {
            top: pos.y + 'px',
            left: (pos.left - 360 - 10) + 'px',
            height: (document.body.clientHeight - pos.y - 60) + 'px',
        }
    }

    return ret
}