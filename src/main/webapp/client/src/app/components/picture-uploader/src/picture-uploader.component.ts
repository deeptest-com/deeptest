import {Component, Input, Output, EventEmitter} from "@angular/core";
import {FileUploader, FileUploaderOptions} from "ng2-file-upload";

import {CONSTANT} from "../../../utils/constant";
import { RouteService } from '../../../service/route';
import {Utils} from "../../../utils/utils";
import { Deferred } from '../../../service/deferred';

@Component({
  selector: 'picture-uploader',
  styleUrls: ['./picture-uploader.scss'],
  templateUrl: './picture-uploader.html',
  providers: []
})
export class PictureUploaderComponent {
  @Output() uploadedEvent: EventEmitter<any> = new EventEmitter<any>();
  @Input() picture: string = '';
  uploader: FileUploader;
  uploadedFile: any;
  hasBaseDropZoneOver: boolean = false;
  public hasAnotherDropZoneOver: boolean = false;

  allowedMimeType: string[] = ['image/png', 'image/jpeg'];
  uploaderOptions: FileUploaderOptions = {
    url: CONSTANT.SERVICE_URL + CONSTANT.UPLOAD_URI,
    authToken: CONSTANT.TOKEN,
    autoUpload: true,
    allowedMimeType: this.allowedMimeType,
    filters: [{
      name: 'upload', fn: (item: any) => {
        return true;
      }
    }]
  };

  constructor(private routeService: RouteService) {
    this.uploader = new FileUploader(this.uploaderOptions);
    this.uploader.onCompleteItem = (item: any, response: any, status: any, headers: any) => {
      this.onUploadCompleteItem(item, response, status, headers);
    };
  }

  public ngOnInit(): void {

  }

  public fileOverBase(e: any): void {
    this.hasBaseDropZoneOver = e;
  }
  public fileOverAnother(e: any): void {
    this.hasAnotherDropZoneOver = e;
  }
  selectFile():void {
    this.uploader.clearQueue();
    jQuery('#upload-input').click();
  }
  fileOver(e:any):void {
    this.hasBaseDropZoneOver = e;
  }
  onUploadCompleteItem (item:any, response:any, status:any, headers:any) {
    let json = JSON.parse(response);
    console.log(json);
    this.uploader.clearQueue();

    if (json.code == 1) {
      this.uploadedFile = json;
      this.picture = json.uploadPath;

      const deferred = new Deferred();
      deferred.promise.then((data) => {
        console.log('onUploadCompleteItem', data);
      }).catch((err) => {console.log('err', err);});

      this.uploadedEvent.emit({
        data: this.picture,
        deferred: deferred,
      });

    } else if (json.code == -100) {
      this.routeService.navTo('/login');
    }
  }

}
