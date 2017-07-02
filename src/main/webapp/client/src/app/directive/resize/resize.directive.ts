import {Directive, ElementRef, Inject, Renderer, OnDestroy, OnInit, AfterViewInit} from "@angular/core";

import * as _ from 'lodash';
declare var jQuery;

@Directive({
  selector: '[resize]'
})
export class ResizeDirective implements OnDestroy, OnInit, AfterViewInit, OnDestroy {

  private elem:Element;

  private container:any;
  private left:any;
  private center:any;
  private right:any;
  private handle1:any;
  private handle2:any;
  private disX:number;

  private isResizing: boolean;
  private lastDownX: any;

  private disposersForDragListeners:Function[] = [];

  public constructor(@Inject(ElementRef) public element:ElementRef,
                     @Inject(Renderer) private renderer:Renderer) {
    this.elem = element.nativeElement;
  }

  public ngOnInit():void {

  }

  ngAfterViewInit() {
    this.container = jQuery(this.elem);

    let left = this.elem.querySelector('.left');
    let center = this.elem.querySelector('.center');
    let right = this.elem.querySelector('.right');
    let handle1 = this.elem.querySelector('.handle1');
    let handle2 = this.elem.querySelector('.handle2');

    this.left = jQuery(left);
    this.center = jQuery(center);
    this.right = jQuery(right);
    this.handle1 = jQuery(handle1);
    this.handle2 = jQuery(handle2);

    this.disposersForDragListeners.push(this.renderer.listen(handle1, 'mousedown', this.onmousedown.bind(this)));
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

    let rightWidth = this.container.width() - (e.clientX - this.container.offset().left);
    this.left.css('right', rightWidth);
    this.left.css('width', e.clientX);
    this.handle1.css('left', e.clientX);
    this.right.css('width', rightWidth);
  }

  private onmouseup(e):any {
    this.isResizing = false;
    _.forEach(this.disposersForDragListeners, (dispose: Function, index: number) => {
        if (index > 0) {
          dispose();
        }
    });
  }
}
