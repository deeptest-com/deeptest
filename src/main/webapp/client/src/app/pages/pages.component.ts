import {Component, ViewEncapsulation} from '@angular/core';
@Component({
  selector: 'pages',
  encapsulation: ViewEncapsulation.None,
  styles: [],
  template: `
    <ba-page-top class="al-header"></ba-page-top>
    
    <div class="al-main">
      <div class="al-content">
        <!--<ba-content-top></ba-content-top>-->
        <router-outlet></router-outlet>
      </div>
    </div>
    
    <footer class="al-footer clearfix">
      <div class="al-footer-right">Created with <i class="ion-heart"></i></div>
      <div class="al-footer-main clearfix">
        <div class="al-copy">&copy; <a href="http://linkr.cn">linkr.cn</a> 2017</div>
        <ul class="al-share clearfix">
          <li><i class="socicon socicon-qq"></i></li>
          <li><i class="socicon socicon-wechat"></i></li>
          <li><i class="socicon socicon-weibo"></i></li>
          <li><i class="socicon socicon-baidu"></i></li>
        </ul>
      </div>
    </footer>
    
    <ba-back-top position="200"></ba-back-top>
    `
})
export class Pages {
  constructor() {

  }

  ngOnInit() {
  }
}
