import * as monaco from "monaco-editor";

export function addExtractAction(editor) {
    console.log('addExtractAction')

    editor.addAction({
        id: 'right-click-text',
        label: '提取为变量',

        keybindings: [
            monaco.KeyMod.CtrlCmd | monaco.KeyCode.F10,
            // chord
            monaco.KeyMod.chord(
                monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyK,
                monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyM
            )
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
        }
    });
}

export function addReplaceAction(editor) {
    console.log('addReplaceAction')

    editor.addAction({
        id: 'right-click-text',
        label: '替换为变量',

        keybindings: [
            monaco.KeyMod.CtrlCmd | monaco.KeyCode.F10,
            // chord
            monaco.KeyMod.chord(
                monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyK,
                monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyM
            )
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
        }
    });
}