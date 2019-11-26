import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(
    private http: HttpClient
  ) {
  }

  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  };

  get() {
    return this.http.get('http://idcards.docker.localhost/api/idCards/')
      .pipe();
  }
  login(payload) {
    return this.http.post('http://admin.docker.localhost/admin/login/', payload)
      .pipe()
  }
}

