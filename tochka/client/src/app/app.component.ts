import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import { NewsItem } from './news-item';
import { debounceTime, distinctUntilChanged, switchMap } from 'rxjs/operators';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  private newsURL = '/news';
  private query = new Subject<string>();
  newsList: Observable<NewsItem[]>;


  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.newsList = this.query.pipe(
      debounceTime(300),
      distinctUntilChanged(),
      switchMap((query: string) => this.getNews(query)),
    );
    this.search('');
  }

  search(query: string): void {
    this.query.next(query.trim());
  }

  getNews(query: string): Observable<NewsItem[]> {
    return this.http.get<NewsItem[]>(`${this.newsURL}?q=${query}`);
  }
}
