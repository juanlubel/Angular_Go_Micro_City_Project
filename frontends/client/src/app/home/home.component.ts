import {Component, OnInit} from '@angular/core';
import {ApiService} from '../core/services';
import {Admin} from '../models';
import {map} from 'rxjs/operators';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  private res: any;


  constructor(
    // private socketService: SocketService,
    private apiService: ApiService
  ) {

  }

  ngOnInit() {

  }

  login(a: string, b: string) {
    const params = {
      name: a,
      pass: b
    };
    this.apiService.login('admin', {
      name: 'admin',
      pass: 'admin'
    })
      .subscribe(
        (res: any) => {
        this.res = res
        console.log('response home component', this.res);
      });
  }
}
