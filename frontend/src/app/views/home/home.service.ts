import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { throwError } from 'rxjs';
import { environment } from '../../../environments/environment';
import { catchError, map } from 'rxjs/operators';
import { ItemProps } from './home.model';

@Injectable()
export class HomeService {
  private readonly API_URL = environment.apiUrl;

  constructor(private http: HttpClient) { }

  getItems(database: string){
    return this.http.get(`${this.API_URL}/${database}`, { observe: 'response' })
      .pipe(map((res :  HttpResponse<any>) => {
        return res.body.items
      }))
      .pipe(catchError(err => {
        return throwError(err);
      })
    );
  }

  insertItem(item: ItemProps, database: string){
    return this.http.post(`${this.API_URL}/${database}`, item, { observe: 'response' })
      .pipe(map((res :  HttpResponse<any>) => {
        return res.body
      }))
      .pipe(catchError(err => {
        return throwError(err);
      })
    );
  }

  deleteItem(item: ItemProps, database: string){
    return this.http.delete(`${this.API_URL}/${database}`, { observe: 'response', body: item })
      .pipe(map((res :  HttpResponse<any>) => {
        return res.body
      }))
      .pipe(catchError(err => {
        return throwError(err);
      })
    );
  }
}
