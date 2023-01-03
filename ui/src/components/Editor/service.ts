import * as monaco from "monaco-editor";

export function addExtractAction(editor, callback) {
    console.log('addExtractAction')

    editor.addAction({
        id: 'extract-action',
        label: '提取变量',

        keybindings: [
            monaco.KeyMod.CtrlCmd | monaco.KeyCode.KEY_R,
            // // chord
            // monaco.KeyMod.chord(
            //     monaco.KeyMod.CtrlCmd | monaco.KeyCode.KEY_K,
            //     monaco.KeyMod.CtrlCmd | monaco.KeyCode.KEY_M
            // )
        ],

        // A precondition for this action.
        precondition: null,

        // A rule to evaluate on top of the precondition in order to dispatch the keybindings.
        keybindingContext: null,
        contextMenuGroupId: 'navigation',
        contextMenuOrder: 1.5,

        run: function (ed) {
            // console.log(ed.getValue());
            console.log(ed.getModel().getValueInRange(ed.getSelection()), ed.getSelection());

            if (ed.getModel().getValueInRange(ed.getSelection()))
                callback({
                    docContent: ed.getValue(),
                    selectContent: ed.getModel().getValueInRange(ed.getSelection()),
                    selectionObj: ed.getSelection()
                })
        }
    });
}

export function addReplaceAction(editor, callback) {
    console.log('addReplaceAction')
    editor.addAction({
        id: 'replace-action',
        label: '使用变量',

        keybindings: [
            monaco.KeyMod.CtrlCmd | monaco.KeyCode.KEY_R,
        ],

        precondition: null,
        keybindingContext: null,
        contextMenuGroupId: 'navigation',
        contextMenuOrder: 1.5,

        run: function (ed) {
            // console.log(ed.getValue());
            console.log(ed.getModel().getValueInRange(ed.getSelection()), ed.getSelection());

            callback({
                docContent: ed.getValue(),
                selectContent: ed.getModel().getValueInRange(ed.getSelection()),
                selectionObj: ed.getSelection()
            })
        }
    });
}