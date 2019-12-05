import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../../environments/environment';
import {Observable, throwError} from 'rxjs';
import {catchError, map} from 'rxjs/operators';

@Injectable()
export class ApiService {
  constructor(
    private http: HttpClient,
  ) {

  }

  private static formatErrors(error: any) {
    console.log(error);
    return throwError(error.error);
  }

  get(path: string) {
    console.log(environment[path]);
    return this.http.get(`${environment[path]}`)
      .pipe(catchError(ApiService.formatErrors));
  }
  // tslint:disable-next-line:ban-types
  post(path: string, body: Object = {}): Observable<any> {
    return this.http.post(`${environment[path]}`, JSON.stringify(body))
      .pipe(catchError(ApiService.formatErrors));
  }
  // tslint:disable-next-line:ban-types
  login(path: string, body: Object = {}): Observable<any>  {
    return this.http.post(`${environment[path]}login`, JSON.stringify(body))
      .pipe(catchError(ApiService.formatErrors));
  }
}
