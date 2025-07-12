import { createClient } from '@supabase/supabase-js';

// // Replace with your actual Supabase Project URL and anon key
const supabaseUrl = import.meta.env.VITE_SUPERBASE_URL;
const supabaseAnonKey = import.meta.env.VITE_SUPERBASE_PUBLIC_KEY;




export const supabase = createClient(supabaseUrl, supabaseAnonKey);
