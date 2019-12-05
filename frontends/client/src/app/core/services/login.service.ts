import {Injectable} from '@angular/core';
import {HttpHeaders} from '@angular/common/http';
import {ApiService} from './api.service';
import {SocketService} from './ws.service';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(

    // private wsService: SocketService
  ) {
  }
  //
  // httpOptions = {
  //   headers: new HttpHeaders({
  //     'Content-Type': 'application/json'
  //   })
  // }
  // login(params: object) {
  //   console.log(params)
  //   const response = this.wsService.send('admin');
  //   console.log(response);
  // }

}

