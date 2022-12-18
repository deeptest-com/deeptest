import request from '@/utils/request';
import {LoginParamsType} from "@/views/user/login/data";

export function initIFrame(html) : any {
    const frameElem = document.getElementById('iframe') as HTMLIFrameElement
    const frameDoc = frameElem.contentWindow?.document as Document

    frameElem.src = "about:blank";
    frameDoc.open();
    frameDoc.write(html);
    frameDoc.close();

    return frameDoc
}
