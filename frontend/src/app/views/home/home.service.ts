import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { ItemProps } from './home.model';

@Injectable()
export class HomeService {
  private readonly API_URL = environment.apiUrl;

  constructor(private http: HttpClient) { }

  getItems(): Observable<ItemProps[]> {
    return this.http.get<ItemProps[]>(`${this.API_URL}/items`);
  }

  insertItem(item: ItemProps){
    return this.http.post(`${this.API_URL}/item`, item);
  }

  deleteItem(id: number){
    return this.http.delete(`${this.API_URL}/item`, { observe: 'response', body: { id } });
  }
}
