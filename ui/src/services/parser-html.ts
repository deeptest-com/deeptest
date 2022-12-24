import request from '@/utils/request';
import {LoginParamsType} from "@/views/user/login/data";
import {Elements} from "@/services/xpath";

export const DeepestKey = 'data-com_deeptest_selection'


export function initIFrame(html) : any {
    const frameElem = document.getElementById('iframe') as HTMLIFrameElement
    const frameDoc = frameElem.contentWindow?.document as Document

    frameElem.src = "about:blank";
    frameDoc.open();
    frameDoc.write(html);
    frameDoc.close();

    return frameDoc
}

export function getXpath(frameDoc: Document) {
    const elem = frameDoc.querySelector(`*[${DeepestKey}="true"]`)

    const xpath = Elements.DOMPath.xPath(elem, true)

    return xpath
}

export function updateElem(frameDoc, docHtml, selectContent, selectionObj) {
    const startLineNumber = selectionObj.startLineNumber - 1
    const endLineNumber = selectionObj.endLineNumber - 1

    const startColumn = selectionObj.startColumn - 1
    const endColumn = selectionObj.endColumn - 1

    const lines = docHtml.split('\n')

    const selectionType = getSelectionType(lines, startLineNumber, endLineNumber, startColumn, endColumn)

    if (selectionType === 'elem') {
        lines[startLineNumber] = lines[startLineNumber].replace(selectContent, selectContent
            + ` ${DeepestKey}="true"`)

        return lines.join('\n')

    } else if (selectionType === 'prop') {
        console.log('')
    } else if (selectionType === 'content') {
        console.log('')
    }

    return
}

export function getSelectionType(lines, startLineNumber, endLineNumber, startColumn, endColumn) {
    const leftNoSpaceChar = getLeftNoSpaceChar(lines, startLineNumber, startColumn)
    const rightChar = getRightChar(lines, endLineNumber, endColumn)

    if (leftNoSpaceChar === '<' && (rightChar === ' ' || rightChar === '>')) {
        return 'elem'
    }

    const leftChar = getLeftChar(lines, startLineNumber, startColumn)
    const rightNoSpaceChar = getRightNoSpaceChar(lines, endLineNumber, endColumn)

    if (leftChar === ' ' && rightNoSpaceChar === '=') {
        return 'prop'
    }

    return 'content'
}

export function getLeftNoSpaceChar(lines, startLineNumber, startColumn) {
    const line = lines[startLineNumber]
    if (startLineNumber === 0 && startColumn === 0) return null

    let leftOne = ''

    if (startColumn > 0) {
        leftOne = line[startColumn-1]
        if (isNotSpace(leftOne)) {
            return leftOne
        }
    }

    startLineNumber -= 1
    startColumn = lines[startLineNumber].length
    if (startLineNumber < 0) return null

    return getLeftNoSpaceChar(lines, startLineNumber, startColumn)
}

export function getRightNoSpaceChar(lines, endLineNumber, endColumn) {
    const line = lines[endLineNumber]
    if (endLineNumber === lines.length - 1 && endColumn === line.length -1) return null

    let rightOne = ''

    if (line.length > endColumn) {
        rightOne = line[endColumn]
        if (rightOne.replace(/(^\s*)|(\s*$)/g, "").length > 0) {
            return rightOne
        }
    }

    endLineNumber += 1
    endColumn = -1
    if (endLineNumber >= lines.length) return null

    return getLeftNoSpaceChar(lines, endLineNumber, endColumn)
}

export function getLeftChar(lines, startLineNumber, startColumn) {
    if (startLineNumber === 0 && startColumn === 0) return null

    const line = lines[startLineNumber]

    if (startColumn > 1) {
        return line[startColumn-1]
    }

    startLineNumber -= 1
    startColumn = lines[startLineNumber].length
    if (startLineNumber < 0) return null

    return getLeftChar(lines, startLineNumber, startColumn)
}

export function getRightChar(lines, endLineNumber, endColumn) {
    const line = lines[endLineNumber]
    if (endLineNumber === lines.length - 1 && endColumn === line.length -1) return null

    if (line.length > endColumn) {
        return line[endColumn]
    }

    endLineNumber += 1
    endColumn = -1
    if (endLineNumber >= lines.length) return null

    return getRightChar(lines, endLineNumber, endColumn)
}

export function findElemLeftStart(lines, startLineNumber, startColumn) { // <

    return
}

export function findElemRightEnd(lines, startLineNumber, startColumn) { // >

    return
}

export function getElemHtml(frameDoc, lines, selection) {


    return
}

export function addPropToElem(elemHtml) {


    return
}

export function isNotSpace(str) {
    return str.replace(/(^\s*)|(\s*$)/g, "").length > 0
}