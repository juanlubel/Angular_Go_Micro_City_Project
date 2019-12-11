import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
  
import { ApiService } from './api.service';
import { Topic } from '..';
import { map, catchError } from 'rxjs/operators';

  @Injectable()
  export class ForumService {
    constructor (
      private apiService: ApiService
    ) {}
    private formatErrors(error: any) {
      return  throwError(error.error);
    }
    getAllTopics(): Observable<[Topic]> {
      return this.apiService.get('topics',"forum_url")
            .pipe(map(data => data.topics));
    }
    createTopic (topic:Topic): Observable<Topic> {
      console.log(topic)
      return this.apiService.post('topics','forum_url', topic)
        .pipe(
          catchError(this.formatErrors)
        );
    }
    deleteTopic (bank:Topic): Observable<Topic> {
      console.log(bank)
      return this.apiService.post('banks','forum_url', bank)
        .pipe(
          catchError(this.formatErrors)
        );
    } 
  }