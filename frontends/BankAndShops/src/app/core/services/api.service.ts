import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';

import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http'; 

import { Observable ,  throwError } from 'rxjs';

/* import { JwtService } from './jwt.service'; */
import { catchError } from 'rxjs/operators';

@Injectable()
export class ApiService {
  constructor(
    private http: HttpClient,
   /*  private jwtService: JwtService */
  ) {}

 

  private formatErrors(error: any) {
    return  throwError(error.error);
  }

  get(path: string,enviroment_constant:string, params: HttpParams = new HttpParams()): Observable<any> {
    return this.http.get(`${environment[enviroment_constant]}${path}`, { params })
      .pipe(catchError(this.formatErrors));
  }

  put(path: string,enviroment_constant:string,  body: Object = {}): Observable<any> {
    return this.http.put(
      `${environment[enviroment_constant]}${path}`,
      JSON.stringify(body)
    ).pipe(catchError(this.formatErrors));
  }

  post(path: string,enviroment_constant:string,  body: Object = {}): Observable<any> {
    return this.http.post(
      `${environment[enviroment_constant]}${path}`,
      JSON.stringify(body)
    ).pipe(catchError(this.formatErrors));
  }

  delete(path,enviroment_constant:string, ): Observable<any> {
    return this.http.delete(
      `${environment[enviroment_constant]}${path}`
    ).pipe(catchError(this.formatErrors));
  }
}
