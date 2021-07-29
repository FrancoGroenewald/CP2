import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { HeadersService } from './headers.service';

@Injectable({
  providedIn: 'root'
})
export class LoginService {
  public apiUrl: string = environment.apiUrl;
  public headerS: any;
  constructor(private http: HttpClient, private headerService: HeadersService) {
    this.headerS = this.headerService.getHTTPHeaders()
  }

  

  public login = (loginValues: any):Observable<any> => {
    return this.http.get(`${this.apiUrl}login`, {params:loginValues, withCredentials:true})
  }
}
