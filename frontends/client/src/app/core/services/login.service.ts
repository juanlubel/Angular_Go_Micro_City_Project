import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {ApiService} from './api.service';
import {Observable} from 'rxjs';
import {Admin} from '../../models';
import {map} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(
    private apiService: ApiService
  ) {
  }

  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  }
  login(params: object): Observable<Admin> {
    return this.apiService.login('admin', params).pipe();

  }
  get() {
    return this.apiService.get('admin').pipe();
  }
}

