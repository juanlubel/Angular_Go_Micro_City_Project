import { Injectable } from '@angular/core';

import { Resolve, Router, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';

import { Topic, ForumService} from '../core';

import { Observable } from 'rxjs';

import { catchError } from 'rxjs/operators';

@Injectable()
export class TopicResolver implements Resolve<Topic> {
  constructor(
    private router: Router,
    private ForumService: ForumService
  ) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> {

    return this.ForumService.getAllTopics().pipe(catchError((err) => this.router.navigateByUrl('/')));

  }
}


