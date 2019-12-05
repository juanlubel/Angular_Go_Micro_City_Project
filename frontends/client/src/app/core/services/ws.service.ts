import { Injectable, EventEmitter } from '@angular/core';

@Injectable()
export class SocketService {

  private socket: WebSocket;
  private listener: EventEmitter<any> = new EventEmitter();

  public constructor() {
    // this.socket = new WebSocket('ws://localhost:3310/ws/login');
    // this.socket.onopen = event => {
    //   this.listener.emit({type: 'open', data: event});
    // };
    // this.socket.onclose = event => {
    //   this.listener.emit({type: 'close', data: event});
    // };
  }

  public send(data: any) {
    this.socket.send(data);
  }

  public close() {
    this.socket.close();
  }

  public getEventListener() {
    return this.listener;
  }

}
