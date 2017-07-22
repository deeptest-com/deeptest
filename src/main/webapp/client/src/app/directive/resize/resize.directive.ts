import {Directive, ElementRef, Inject, Renderer2, OnDestroy, OnInit, AfterViewInit} from "@angular/core";

import * as _ from 'lodash';
declare var jQuery;

import {UserService} from '../../service/user';

@Directive({
  selector: '[resize]'
})
export class ResizeDirective implements OnDestroy, OnInit, AfterViewInit, OnDestroy {

  private elem:Element;

  private container:any;
  private left:any;
  private center:any;
  private right:any;
  private handleId:any;
  private handle1:any;
  private handle2:any;
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
    let center = this.elem.querySelector('.resize-center');
    let right = this.elem.querySelector('.resize-right');
    let handle1 = this.elem.querySelector('.resize-handle1');
    let handle2 = this.elem.querySelector('.resize-handle2');

    this.left = jQuery(left);
    this.center = jQuery(center);
    this.right = jQuery(right);
    this.handle1 = jQuery(handle1);
    this.handle2 = jQuery(handle2);

    this.disposersForDragListeners.push(this.renderer.listen(this.elem, 'mousedown', this.onmousedown.bind(this)));
  }

  private onmousedown(e):any {
    this.handleId = e.target.id;

    if (this.handleId != 'handle1' && this.handleId != 'handle2') {
      return;
    }

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

    let containerWidth = this.container.width();
    if (this.handleId === 'handle1') {

      let rightOrigWidth = parseInt(this.right.css('width'));
      let centerWidth = containerWidth - e.clientX - rightOrigWidth;

      this.handle1.css('left', e.clientX);

      this.left.css('width', e.clientX);
      this.center.css('left', e.clientX);
      this.center.css('width', centerWidth);
    } else if (this.handleId === 'handle2') {

      let leftOrigWidth = parseInt(this.left.css('width'));
      let rightWidth = containerWidth - e.clientX;
      let centerWidth = e.clientX - leftOrigWidth;

      this.handle2.css('left', e.clientX);

      this.right.css('left', e.centerWidth);
      this.right.css('width', rightWidth);
      this.center.css('width', centerWidth);
    }
  }

  private onmouseup(e):any {
    this.handleId = undefined;

    this.userService.setSize(parseInt(this.left.css('width')), parseInt(this.right.css('width'))).subscribe((json:any) => {
      console.log(json.code);
    });

    this.isResizing = false;
    _.forEach(this.disposersForDragListeners, (dispose: Function, index: number) => {
        if (index > 0) {
          dispose();
        }
    });
  }

  public ngOnDestroy():void {
    this.disposersForDragListeners.forEach(dispose => dispose());
  }
}
