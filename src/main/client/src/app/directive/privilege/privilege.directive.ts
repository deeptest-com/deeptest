import {Directive, ElementRef, Inject, Renderer, Input, OnInit, OnDestroy} from "@angular/core";

import {CONSTANT} from '../../utils/constant';

import {GlobalState} from "../../global.state";
import { PrivilegeService } from '../../service/privilege';

@Directive({
  selector: '[privilege]'
})
export class PrivilegeDirective implements OnInit, OnDestroy {
  eventCode:string = 'PrivilegeDirective';

  private elem:Element;
  @Input() privs: string;

  public constructor(private _state: GlobalState, private _privilegeService: PrivilegeService,
                     @Inject(ElementRef) public element:ElementRef, @Inject(Renderer) private renderer:Renderer) {
    this.elem = element.nativeElement;
  }

  public ngOnInit():void {
      this.update();
  }

  update():void {
    let ret = this._privilegeService.hasPrivilege(this.privs);

    if (!ret) {
      this.renderer.setElementStyle(this.elem, 'display', 'none');
    }
  }

  ngOnDestroy(): void {

  };
}
