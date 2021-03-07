export type SearchedObject = {
    objectID: number|null
    img: string //base64
    title: string
    description: string
};

export type MultipleSearchedObjects = SearchedObject[];

export type ObjectVendor = {
    vendorID: number
    shopLogo: Blob
    stars: 0|0.5|1|1.5|2|2.5|3|3.5|4|4.5|5
    opinions: number
    price: number
    url: string
};

export type SearchedSaleOffer = {
    vendor: ObjectVendor
    object: SearchedObject
};

export type MultipleObjectVendors = ObjectVendor[];

export type VendorObjectHistory = {
    vendorID: number
    data: {
        value: number
        date: Date
    }[];
}

export type SearchedObjectHistory = {
    object: SearchedObject
    history: VendorObjectHistory[]
};