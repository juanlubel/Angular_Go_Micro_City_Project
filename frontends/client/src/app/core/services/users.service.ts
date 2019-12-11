import {Injectable} from '@angular/core';

import {HttpHeaders} from '@angular/common/http';
import {ApiService} from './api.service';


@Injectable({
  providedIn: 'root'
})
export class UsersService {

  constructor(
    private apiService: ApiService
  ) {
  }

  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  };
  getUsers() {
    return this.apiService.get('user').pipe();

  }
  // login() {
  //   return this.apiService.login('idCards', {
  //     name: 'admin',
  //     pass: 'admin'
  //   }).pipe();
  // }
}
