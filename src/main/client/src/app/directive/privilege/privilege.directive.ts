import {Directive, ElementRef, Inject, Renderer, Input, OnInit} from "@angular/core";

import {CONSTANT} from '../../utils/constant';

import {GlobalState} from "../../global.state";

@Directive({
  selector: '[privilege]'
})
export class PrivilegeDirective implements OnInit {

  private elem:Element;
  @Input() privs: string;

  public constructor(private _state: GlobalState, @Inject(ElementRef) public element:ElementRef, @Inject(Renderer) private renderer:Renderer) {
    this.elem = element.nativeElement;
  }

  public ngOnInit():void {
    if (!CONSTANT.PROFILE || !CONSTANT.PROFILE.projectPrivilege) {
      this._state.subscribe(CONSTANT.STATE_CHANGE_PROFILE, (profile) => {
        console.log(CONSTANT.STATE_CHANGE_PROFILE + ' in PrivilegeDirective', profile);
        this.update();
      });
      console.log('===aaa', this.privs);
    } else {
      this.update();
      console.log('===bbb', this.privs);
    }
  }

  update():void {
    let ret = true;

    let arr = this.privs.split(',');
    for (let i = 0; i < arr.length; i++) {
      if (!CONSTANT.PROFILE.projectPrivilege[arr[i]]) {
        ret = false;
      }
    }

    if (!ret) {
      this.renderer.setElementStyle(this.elem, 'display', 'none');
    }
  }

}
