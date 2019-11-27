import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {environment} from '../../environments/environment';
import {catchError} from 'rxjs/operators';


@Injectable({
  providedIn: 'root'
})
export class ApiService {
  constructor(
    private http: HttpClient,
  ) {}

  get(path: string) {
    console.log(`${environment[path]}${path}/`);
    return this.http.get(`${environment[path]}`)
      .pipe();
  }
  // tslint:disable-next-line:ban-types
  post(path: string, payload: Object = {}) {
    return this.http.post(`${environment[path]}`, JSON.stringify(payload))
      .pipe();
  }
  // tslint:disable-next-line:ban-types
  login(path: string, payload: Object = {}) {
    return this.http.post(`${environment[path]}/login`, JSON.stringify(payload))
      .pipe();
  }
}
