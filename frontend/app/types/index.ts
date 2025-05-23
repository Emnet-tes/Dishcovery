export interface Recipe {
  id: number;
  title: string;
  author: string;
  rating: number;
  reviews: number;
  time: string;
  timeRange?: string; // Optional, based on previous work
  image: string;
  // Add other potential recipe properties as needed
} 