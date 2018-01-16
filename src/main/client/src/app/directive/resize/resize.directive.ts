import {Directive, ElementRef, Inject, Renderer2, OnDestroy, OnInit, AfterViewInit} from "@angular/core";

import * as _ from 'lodash';
declare var jQuery;

import { CONSTANT } from '../../utils/constant';
import { UserService } from '../../service/user';

@Directive({
  selector: '[resize]'
})
export class ResizeDirective implements OnDestroy, OnInit, AfterViewInit, OnDestroy {

  private elem:Element;

  private container:any;
  private left:any;
  private right:any;
  private handle:any;
  private disX:number;

  private isResizing: boolean;
  private lastDownX: any;

  private disposersForDragListeners:Function[] = [];

  public constructor(@Inject(ElementRef) public element:ElementRef, @Inject(Renderer2) private renderer:Renderer2,
                     private userService: UserService) {
    this.elem = element.nativeElement;
  }

  public ngOnInit():void {

  }

  ngAfterViewInit() {
    this.container = jQuery(this.elem);

    let left = this.elem.querySelector('.resize-left');
    // let right = this.elem.querySelector('.resize-right');
    let handle = this.elem.querySelector('.resize-handle');

    this.left = jQuery(left);
    // this.right = jQuery(right);
    this.handle = jQuery(handle);

    // let rightWidth = this.container.width() - (this.left - this.container.offset().left);
    // this.right.css('width', rightWidth);

    this.disposersForDragListeners.push(
      this.renderer.listen(handle, 'mousedown', this.onmousedown.bind(this)));
  }

  public ngOnDestroy():void {
    this.disposersForDragListeners.forEach(dispose => dispose());
  }

  private onmousedown(e):any {
    this.isResizing = true;
    this.lastDownX = e.clientX;

    this.disposersForDragListeners.push(
      this.renderer.listen(this.elem, 'mousemove', this.onmousemove.bind(this)));
    this.disposersForDragListeners.push(
      this.renderer.listen(this.elem, 'mouseup', this.onmouseup.bind(this)));
  }

  private onmousemove(e):any {
    if (!this.isResizing) {
      return;
    }

    // let rightWidth = this.container.width() - (e.clientX - this.container.offset().left);
    // this.left.css('right', rightWidth);
    this.left.css('width', e.clientX);
    this.handle.css('left', e.clientX);
    // this.right.css('width', rightWidth);
  }

  private onmouseup(e):any {
    this.isResizing = false;
    _.forEach(this.disposersForDragListeners, (dispose: Function, index: number) => {
      if (index > 0) {
        dispose();
      }
    });

    let left = this.left.css('width').replace('px', '');
    this.userService.setLeftSize(left).subscribe((json:any) => {
      CONSTANT.PROFILE.leftSize = left;
    });
  }
}
