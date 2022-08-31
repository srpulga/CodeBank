export interface Product{
  id: string;
  name: string;
  description: string;
  image_url: string;
  slug: string;
  price: number;
  created_at: string;
}

export const products: Product[] = [
  {
    id: 'uuid',
    name: 'Product 1',
    description: 'This is a product',
    price: 10.10,
    image_url: 'https://source.unsplash.com/random?product,' + Math.random(),
    slug: 'product-1',
    created_at: '2022-08-31T00:00:00',
  },
  {
    id: 'uuid',
    name: 'Product 2',
    description: 'This is a product',
    price: 10.10,
    image_url: 'https://source.unsplash.com/random?product,' + Math.random(),
    slug: 'product-2',
    created_at: '2022-08-31T00:00:00',
  },
];
