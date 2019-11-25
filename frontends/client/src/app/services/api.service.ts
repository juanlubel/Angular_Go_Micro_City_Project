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
    return this.http.get(`${environment[path]}`).pipe();
  }
}
