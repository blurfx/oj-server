import { NextResponse } from 'next/server';

export async function POST() {
  try {
    const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/v1/auth/logout`, {
      method: 'POST',
      credentials: 'include',
    });

    if (!response.ok) {
      const error = await response.json();
      return NextResponse.json(
        { error: error.error || 'Logout failed' },
        { status: response.status }
      );
    }

    return NextResponse.json({ success: true });
  } catch (error) {
    return NextResponse.json(
      { error: 'Internal server error' },
      { status: 500 }
    );
  }
} 