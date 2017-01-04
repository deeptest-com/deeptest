import { Inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

declare var BMap: any;

@Injectable()
export class MapService {
    static AK: string = '8YpOIjm9GmSbsONN12tSvb24';
    map: any;

    loadMap(city: string, keyword: string): void {
        let me = this;
        console.log(city + '->' + keyword);

        if (!me.map) {
          me.map = new BMap.Map("map");
        }
        var localSearch = new BMap.LocalSearch(me.map);
        localSearch.setLocation(city)
        localSearch.enableAutoViewport();

        localSearch.setSearchCompleteCallback(function (searchResult) {
          var poi = searchResult.getPoi(0);
          me.map.centerAndZoom(poi.point, 13);

          var marker = new BMap.Marker(poi.point);
          me.map.addOverlay(marker);

          var label = new BMap.Label(keyword, {offset:new BMap.Size(10,-10)});
          marker.setLabel(label);
        });

        localSearch.search(keyword);
    }
}
