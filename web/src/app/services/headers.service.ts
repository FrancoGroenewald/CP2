import { HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class HeadersService {

  constructor() { }

  public getHTTPHeaders = () => {
    const jwt_token = localStorage.getItem('jwtToken');
    if (localStorage.getItem('jwtToken')) {
        const httpHeader = new HttpHeaders({
            'Content-Type': 'application/json',
            'X-Access-Token': jwt_token || ''
        });
        return httpHeader;
    } else {
        console.log('JWT Token not set!');
        return;
    }
  }
}
