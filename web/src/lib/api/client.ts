import { z } from "zod";
import * as api from "./types";
import { BaseApiClient, createUrl, type ExtraOptions } from "./base-client";


export class ApiClient extends BaseApiClient {
  url: ClientUrls;

  constructor(baseUrl: string) {
    super(baseUrl);
    this.url = new ClientUrls(baseUrl);
  }
  
  createCollection(body: api.CreateCollectionBody, options?: ExtraOptions) {
    return this.request("/api/v1/collections", "POST", api.CreateCollection, z.any(), body, options)
  }
  
  deleteCollection(id: string, options?: ExtraOptions) {
    return this.request(`/api/v1/collections/${id}`, "DELETE", z.undefined(), z.any(), undefined, options)
  }
  
  editCollection(id: string, body: api.EditCollectionBody, options?: ExtraOptions) {
    return this.request(`/api/v1/collections/${id}`, "PATCH", z.undefined(), z.any(), body, options)
  }
  
  getCollectionById(id: string, options?: ExtraOptions) {
    return this.request(`/api/v1/collections/${id}`, "GET", api.GetCollectionById, z.any(), undefined, options)
  }
  
  
  getCollectionImages(id: string, options?: ExtraOptions) {
    return this.request(`/api/v1/collections/${id}/images`, "GET", api.GetCollectionImages, z.any(), undefined, options)
  }
  
  getCollections(options?: ExtraOptions) {
    return this.request("/api/v1/collections", "GET", api.GetCollection, z.any(), undefined, options)
  }
  
  getSystemInfo(options?: ExtraOptions) {
    return this.request("/api/v1/system/info", "GET", api.GetSystemInfo, z.any(), undefined, options)
  }
  
  signin(body: api.SigninBody, options?: ExtraOptions) {
    return this.request("/api/v1/auth/signin", "POST", api.Signin, z.any(), body, options)
  }
  
  uploadToCollection(id: string, body: FormData, options?: ExtraOptions) {
    return this.requestForm(`/api/v1/collections/${id}/upload`, "POST", z.undefined(), z.any(), body, options)
  }
}

export class ClientUrls {
  baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }
  
  createCollection() {
    return createUrl(this.baseUrl, "/api/v1/collections")
  }
  
  deleteCollection(id: string) {
    return createUrl(this.baseUrl, `/api/v1/collections/${id}`)
  }
  
  editCollection(id: string) {
    return createUrl(this.baseUrl, `/api/v1/collections/${id}`)
  }
  
  getCollectionById(id: string) {
    return createUrl(this.baseUrl, `/api/v1/collections/${id}`)
  }
  
  getCollectionImage(id: string, file: string) {
    return createUrl(this.baseUrl, `/files/collections/${id}/images/${file}`)
  }
  
  getCollectionImages(id: string) {
    return createUrl(this.baseUrl, `/api/v1/collections/${id}/images`)
  }
  
  getCollections() {
    return createUrl(this.baseUrl, "/api/v1/collections")
  }
  
  getSystemInfo() {
    return createUrl(this.baseUrl, "/api/v1/system/info")
  }
  
  signin() {
    return createUrl(this.baseUrl, "/api/v1/auth/signin")
  }
  
  uploadToCollection(id: string) {
    return createUrl(this.baseUrl, `/api/v1/collections/${id}/upload`)
  }
}
