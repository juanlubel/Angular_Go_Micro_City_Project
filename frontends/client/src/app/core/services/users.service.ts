import {Injectable} from '@angular/core';

import {HttpHeaders} from '@angular/common/http';
import {ApiService} from './api.service';
import {User} from '../../models';


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
    return this.apiService.get('users').pipe();
  }
  updateUser(user: User) {
    return this.apiService.put('user', user).pipe();
  }
  // login() {
  //   return this.apiService.login('idCards', {
  //     name: 'admin',
  //     pass: 'admin'
  //   }).pipe();
  // }
}
