import {Injectable, OnInit} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {environment} from '../../../environments/environment';
import {Observable, throwError} from 'rxjs';
import {catchError} from 'rxjs/operators';
import {TokenType, User} from '../../models';
import {Store} from '@ngrx/store';
import {IAppState} from '../../store/state/app.state';
import {selectUserLogged} from '../../store/selectors/user.selector';
import {GetToken} from '../../store/actions/user.actions';

@Injectable()
export class ApiService implements OnInit {
  public token$ = new Observable<string>();
  name: TokenType;
  constructor(
    private http: HttpClient,
    private store: Store<IAppState>
  ) {
    this.token$ = this.store.select(selectUserLogged);
    const name = localStorage.getItem('userLoggedName');

    const headers: HttpHeaders = new HttpHeaders({bearer: ' this.token$'});
  }
  private static formatErrors(error: any) {
    console.log(error);
    return throwError(error.error);
  }

  // tslint:disable-next-line:contextual-lifecycle
  ngOnInit() {
    this.store.dispatch(new GetToken(this.name));
  }

  get(path: string, type: string  = '', id: string = '') {
    console.log(environment[path]);
    return this.http.get(`${environment[path]}${type}${id}`)
      .pipe(catchError(ApiService.formatErrors));
  }
  // tslint:disable-next-line:ban-types
  post(path: string, body: Object = {}): Observable<any> {
    return this.http.post(`${environment[path]}`, JSON.stringify(body))
      .pipe(catchError(ApiService.formatErrors));
  }
  // tslint:disable-next-line:ban-types
  put(path: string, body: User): Observable<any> {
    console.log(`${environment[path]}${body.slug}`);
    return this.http.put(`${environment[path]}${body.slug}`, JSON.stringify(body))
      .pipe(catchError(ApiService.formatErrors));
  }
  // tslint:disable-next-line:ban-types
  login(path: string, body: Object = {}): Observable<any>  {
    return this.http.post(`${environment[path]}login`, JSON.stringify(body))
      .pipe(catchError(ApiService.formatErrors));
  }
}
