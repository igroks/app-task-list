import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { throwError } from 'rxjs';
import { environment } from '../../../environments/environment';
import { catchError, map } from 'rxjs/operators';
import { ItemProps } from './home.model';

@Injectable()
export class HomeService {
  private readonly API_URL = environment.apiUrl + '/task';

  constructor(private http: HttpClient) { }

  getItems(){
    return this.http.get(`${this.API_URL}`, { observe: 'response' })
      .pipe(map((res :  HttpResponse<any>) => {
        return res.body.items
      }))
      .pipe(catchError(err => {
        return throwError(err);
      })
    );
  }

  insertItem(item: ItemProps){
    return this.http.post(`${this.API_URL}`, item, { observe: 'response' })
      .pipe(map((res :  HttpResponse<any>) => {
        return res.body
      }))
      .pipe(catchError(err => {
        return throwError(err);
      })
    );
  }

  deleteItem(id: number){
    return this.http.delete(`${this.API_URL}/${id}`, { observe: 'response' })
      .pipe(map((res :  HttpResponse<any>) => {
        return res.body
      }))
      .pipe(catchError(err => {
        return throwError(err);
      })
    );
  }
}
