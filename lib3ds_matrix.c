/*
    Copyright (C) 1996-2008 by Jan Eric Kyprianidis <www.kyprianidis.com>
    All rights reserved.
    
    This program is free  software: you can redistribute it and/or modify 
    it under the terms of the GNU Lesser General Public License as published 
    by the Free Software Foundation, either version 2.1 of the License, or 
    (at your option) any later version.

    Thisprogram  is  distributed in the hope that it will be useful, 
    but WITHOUT ANY WARRANTY; without even the implied warranty of 
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
    GNU Lesser General Public License for more details.
    
    You should  have received a copy of the GNU Lesser General Public License
    along with  this program; If not, see <http://www.gnu.org/licenses/>. 
*/

/** @file lib3ds_matrix.c
	Matrix mathematics implementation */

#include "lib3ds_impl.h"


/*!
 * Clear a matrix to all zeros.
 *
 * \param m Matrix to be cleared.
 */
void
lib3ds_matrix_zero(float m[4][4]) {
    int i, j;

    for (i = 0; i < 4; i++) {
        for (j = 0; j < 4; j++) m[i][j] = 0.0f;
    }
}


/*!
 * Set a matrix to identity.
 *
 * \param m Matrix to be set.
 */
void
lib3ds_matrix_identity(float m[4][4]) {
    int i, j;

    for (i = 0; i < 4; i++) {
        for (j = 0; j < 4; j++) m[i][j] = 0.0;
    }
    for (i = 0; i < 4; i++) m[i][i] = 1.0;
}


/*!
 * Copy a matrix.
 */
void
lib3ds_matrix_copy(float dest[4][4], float src[4][4]) {
    memcpy(dest, src, 16 * sizeof(float));
}

void
lib3ds_matrix_copy_to_double(double dest[4][4], float src[4][4]) {
  for(int i = 0;i<4;i++){
    for(int j = 0;j<4;j++){
      dest[i][j] = src[i][j];
    }
  }
}



/*!
 * Negate a matrix -- all elements negated.
 */
void
lib3ds_matrix_neg(float m[4][4]) {
    int i, j;

    for (j = 0; j < 4; j++) {
        for (i = 0; i < 4; i++) {
            m[j][i] = -m[j][i];
        }
    }
}


/*!
 * Transpose a matrix in place.
 */
void
lib3ds_matrix_transpose(float m[4][4]) {
    int i, j;
    float swp;

    for (j = 0; j < 4; j++) {
        for (i = j + 1; i < 4; i++) {
            swp = m[j][i];
            m[j][i] = m[i][j];
            m[i][j] = swp;
        }
    }
}


/*!
 * Add two matrices.
 */
void
lib3ds_matrix_add(float m[4][4], float a[4][4], float b[4][4]) {
    int i, j;

    for (j = 0; j < 4; j++) {
        for (i = 0; i < 4; i++) {
            m[j][i] = a[j][i] + b[j][i];
        }
    }
}


/*!
 * Subtract two matrices.
 *
 * \param m Result.
 * \param a Addend.
 * \param b Minuend.
 */
void
lib3ds_matrix_sub(float m[4][4], float a[4][4], float b[4][4]) {
    int i, j;

    for (j = 0; j < 4; j++) {
        for (i = 0; i < 4; i++) {
            m[j][i] = a[j][i] - b[j][i];
        }
    }
}


/*!
 * Multiplies a matrix by a second one (m = m * n).
 */
void
lib3ds_matrix_mult(float m[4][4], float a[4][4], float b[4][4]) {
    float tmp[4][4];
    int i, j, k;
    float ab;

    memcpy(tmp, a, 16 * sizeof(float));
    for (j = 0; j < 4; j++) {
        for (i = 0; i < 4; i++) {
            ab = 0.0f;
            for (k = 0; k < 4; k++) ab += tmp[k][i] * b[j][k];
            m[j][i] = ab;
        }
    }
}

void lib3ds_matrix_mult_double(double m[4][4], double a[4][4], double b[4][4]) {
  double tmp[4][4];
  int i, j, k;
  double ab;
  
  memcpy(tmp, a, 16 * sizeof(double));
  for (j = 0; j < 4; j++) {
    for (i = 0; i < 4; i++) {
      ab = 0.0f;
      for (k = 0; k < 4; k++) ab += tmp[k][i] * b[j][k];
      m[j][i] = ab;
    }
  }
}



/*!
 * Multiply a matrix by a scalar.
 *
 * \param m Matrix to be set.
 * \param k Scalar.
 */
void
lib3ds_matrix_scalar(float m[4][4], float k) {
    int i, j;

    for (j = 0; j < 4; j++) {
        for (i = 0; i < 4; i++) {
            m[j][i] *= k;
        }
    }
}


static float
det2x2(
    float a, float b,
    float c, float d) {
    return((a)*(d) - (b)*(c));
}


static float
det3x3(
    float a1, float a2, float a3,
    float b1, float b2, float b3,
    float c1, float c2, float c3) {
    return(
              a1*det2x2(b2, b3, c2, c3) -
              b1*det2x2(a2, a3, c2, c3) +
              c1*det2x2(a2, a3, b2, b3)
          );
}


/*!
 * Find determinant of a matrix.
 */
float
lib3ds_matrix_det(float m[4][4]) {
    float a1, a2, a3, a4, b1, b2, b3, b4, c1, c2, c3, c4, d1, d2, d3, d4;

    a1 = m[0][0];
    b1 = m[1][0];
    c1 = m[2][0];
    d1 = m[3][0];
    a2 = m[0][1];
    b2 = m[1][1];
    c2 = m[2][1];
    d2 = m[3][1];
    a3 = m[0][2];
    b3 = m[1][2];
    c3 = m[2][2];
    d3 = m[3][2];
    a4 = m[0][3];
    b4 = m[1][3];
    c4 = m[2][3];
    d4 = m[3][3];
    return(
              a1 * det3x3(b2, b3, b4, c2, c3, c4, d2, d3, d4) -
              b1 * det3x3(a2, a3, a4, c2, c3, c4, d2, d3, d4) +
              c1 * det3x3(a2, a3, a4, b2, b3, b4, d2, d3, d4) -
              d1 * det3x3(a2, a3, a4, b2, b3, b4, c2, c3, c4)
          );
}


/*!
 * Invert a matrix in place.
 *
 * \param m Matrix to invert.
 *
 * \return LIB3DS_TRUE on success, LIB3DS_FALSE on failure.
 *
 * GGemsII, K.Wu, Fast Matrix Inversion
 */
int
lib3ds_matrix_inv(float m[4][4]) {
    int i, j, k;
    int pvt_i[4], pvt_j[4];            /* Locations of pivot elements */
    float pvt_val;               /* Value of current pivot element */
    float hold;                  /* Temporary storage */
    float determinat;

    determinat = 1.0f;
    for (k = 0; k < 4; k++)  {
        /* Locate k'th pivot element */
        pvt_val = m[k][k];          /* Initialize for search */
        pvt_i[k] = k;
        pvt_j[k] = k;
        for (i = k; i < 4; i++) {
            for (j = k; j < 4; j++) {
                if (fabs(m[i][j]) > fabs(pvt_val)) {
                    pvt_i[k] = i;
                    pvt_j[k] = j;
                    pvt_val = m[i][j];
                }
            }
        }

        /* Product of pivots, gives determinant when finished */
        determinat *= pvt_val;
        if (fabs(determinat) < LIB3DS_EPSILON) {
            return(FALSE);  /* Matrix is singular (zero determinant) */
        }

        /* "Interchange" rows (with sign change stuff) */
        i = pvt_i[k];
        if (i != k) {             /* If rows are different */
            for (j = 0; j < 4; j++) {
                hold = -m[k][j];
                m[k][j] = m[i][j];
                m[i][j] = hold;
            }
        }

        /* "Interchange" columns */
        j = pvt_j[k];
        if (j != k) {            /* If columns are different */
            for (i = 0; i < 4; i++) {
                hold = -m[i][k];
                m[i][k] = m[i][j];
                m[i][j] = hold;
            }
        }

        /* Divide column by minus pivot value */
        for (i = 0; i < 4; i++) {
            if (i != k) m[i][k] /= (-pvt_val) ;
        }

        /* Reduce the matrix */
        for (i = 0; i < 4; i++) {
            hold = m[i][k];
            for (j = 0; j < 4; j++) {
                if (i != k && j != k) m[i][j] += hold * m[k][j];
            }
        }

        /* Divide row by pivot */
        for (j = 0; j < 4; j++) {
            if (j != k) m[k][j] /= pvt_val;
        }

        /* Replace pivot by reciprocal (at last we can touch it). */
        m[k][k] = 1.0f / pvt_val;
    }

    /* That was most of the work, one final pass of row/column interchange */
    /* to finish */
    for (k = 4 - 2; k >= 0; k--) { /* Don't need to work with 1 by 1 corner*/
        i = pvt_j[k];          /* Rows to swap correspond to pivot COLUMN */
        if (i != k) {          /* If rows are different */
            for (j = 0; j < 4; j++) {
                hold = m[k][j];
                m[k][j] = -m[i][j];
                m[i][j] = hold;
            }
        }

        j = pvt_i[k];         /* Columns to swap correspond to pivot ROW */
        if (j != k)           /* If columns are different */
            for (i = 0; i < 4; i++) {
                hold = m[i][k];
                m[i][k] = -m[i][j];
                m[i][j] = hold;
            }
    }
    return(TRUE);
}

int
lib3ds_matrix_inv_double(double m[4][4]) {
  int i, j, k;
  int pvt_i[4], pvt_j[4];            /* Locations of pivot elements */
  double pvt_val;               /* Value of current pivot element */
  double hold;                  /* Temporary storage */
  double determinat;
  
  determinat = 1.0f;
  for (k = 0; k < 4; k++)  {
    /* Locate k'th pivot element */
    pvt_val = m[k][k];          /* Initialize for search */
    pvt_i[k] = k;
    pvt_j[k] = k;
    for (i = k; i < 4; i++) {
      for (j = k; j < 4; j++) {
        if (fabs(m[i][j]) > fabs(pvt_val)) {
          pvt_i[k] = i;
          pvt_j[k] = j;
          pvt_val = m[i][j];
        }
      }
    }
    
    /* Product of pivots, gives determinant when finished */
    determinat *= pvt_val;
    if (fabs(determinat) < LIB3DS_EPSILON) {
      return(FALSE);  /* Matrix is singular (zero determinant) */
    }
    
    /* "Interchange" rows (with sign change stuff) */
    i = pvt_i[k];
    if (i != k) {             /* If rows are different */
      for (j = 0; j < 4; j++) {
        hold = -m[k][j];
        m[k][j] = m[i][j];
        m[i][j] = hold;
      }
    }
    
    /* "Interchange" columns */
    j = pvt_j[k];
    if (j != k) {            /* If columns are different */
      for (i = 0; i < 4; i++) {
        hold = -m[i][k];
        m[i][k] = m[i][j];
        m[i][j] = hold;
      }
    }
    
    /* Divide column by minus pivot value */
    for (i = 0; i < 4; i++) {
      if (i != k) m[i][k] /= (-pvt_val) ;
    }
    
    /* Reduce the matrix */
    for (i = 0; i < 4; i++) {
      hold = m[i][k];
      for (j = 0; j < 4; j++) {
        if (i != k && j != k) m[i][j] += hold * m[k][j];
      }
    }
    
    /* Divide row by pivot */
    for (j = 0; j < 4; j++) {
      if (j != k) m[k][j] /= pvt_val;
    }
    
    /* Replace pivot by reciprocal (at last we can touch it). */
    m[k][k] = 1.0f / pvt_val;
  }
  
  /* That was most of the work, one final pass of row/column interchange */
  /* to finish */
  for (k = 4 - 2; k >= 0; k--) { /* Don't need to work with 1 by 1 corner*/
    i = pvt_j[k];          /* Rows to swap correspond to pivot COLUMN */
    if (i != k) {          /* If rows are different */
      for (j = 0; j < 4; j++) {
        hold = m[k][j];
        m[k][j] = -m[i][j];
        m[i][j] = hold;
      }
    }
    
    j = pvt_i[k];         /* Columns to swap correspond to pivot ROW */
    if (j != k)           /* If columns are different */
      for (i = 0; i < 4; i++) {
        hold = m[i][k];
        m[i][k] = -m[i][j];
        m[i][j] = hold;
      }
  }
  return(TRUE);
}


/*!
 * Apply a translation to a matrix.
 */
void
lib3ds_matrix_translate(float m[4][4], float x, float y, float z) {
    int i;

    for (i = 0; i < 3; i++) {
        m[3][i] += m[0][i] * x + m[1][i] * y + m[2][i] * z;
    }
}

void
lib3ds_matrix_translate_to_double(double m[4][4], float x, float y, float z) {
  int i;
  
  for (i = 0; i < 3; i++) {
    m[3][i] += m[0][i] * x + m[1][i] * y + m[2][i] * z;
  }
}


/*!
 * Apply scale factors to a matrix.
 */
void
lib3ds_matrix_scale(float m[4][4], float x, float y, float z) {
    int i;

    for (i = 0; i < 4; i++) {
        m[0][i] *= x;
        m[1][i] *= y;
        m[2][i] *= z;
    }
}


/*!
 * Apply a rotation about an arbitrary axis to a matrix.
 */
void
lib3ds_matrix_rotate_quat(float m[4][4], float q[4]) {
    float s, xs, ys, zs, wx, wy, wz, xx, xy, xz, yy, yz, zz, l;
    float R[4][4];

    l = q[0] * q[0] + q[1] * q[1] + q[2] * q[2] + q[3] * q[3];
    if (fabs(l) < LIB3DS_EPSILON) {
        s = 1.0f;
    } else {
        s = 2.0f / l;
    }

    xs = q[0] * s;
    ys = q[1] * s;
    zs = q[2] * s;
    wx = q[3] * xs;
    wy = q[3] * ys;
    wz = q[3] * zs;
    xx = q[0] * xs;
    xy = q[0] * ys;
    xz = q[0] * zs;
    yy = q[1] * ys;
    yz = q[1] * zs;
    zz = q[2] * zs;

    R[0][0] = 1.0f - (yy + zz);
    R[1][0] = xy - wz;
    R[2][0] = xz + wy;
    R[0][1] = xy + wz;
    R[1][1] = 1.0f - (xx + zz);
    R[2][1] = yz - wx;
    R[0][2] = xz - wy;
    R[1][2] = yz + wx;
    R[2][2] = 1.0f - (xx + yy);
    R[3][0] = R[3][1] = R[3][2] = R[0][3] = R[1][3] = R[2][3] = 0.0f;
    R[3][3] = 1.0f;

    lib3ds_matrix_mult(m, m, R);
}


/*!
 * Apply a rotation about an arbitrary axis to a matrix.
 */
void
lib3ds_matrix_rotate(float m[4][4], float angle, float ax, float ay, float az) {
    float q[4];
    float axis[3];

    lib3ds_vector_make(axis, ax, ay, az);
    lib3ds_quat_axis_angle(q, axis, angle);
    lib3ds_matrix_rotate_quat(m, q);
}


/*!
 * Compute a camera matrix based on position, target and roll.
 *
 * Generates a translate/rotate matrix that maps world coordinates
 * to camera coordinates.  Resulting matrix does not include perspective
 * transform.
 *
 * \param matrix Destination matrix.
 * \param pos Camera position
 * \param tgt Camera target
 * \param roll Roll angle
 */
void
lib3ds_matrix_camera(float matrix[4][4], float pos[3], float tgt[3], float roll) {
    float M[4][4];
    float x[3], y[3], z[3];

    lib3ds_vector_sub(y, tgt, pos);
    lib3ds_vector_normalize(y);

    if (y[0] != 0. || y[1] != 0) {
        z[0] = 0;
        z[1] = 0;
        z[2] = 1.0;
    } else { /* Special case:  looking straight up or down z axis */
        z[0] = -1.0;
        z[1] = 0;
        z[2] = 0;
    }

    lib3ds_vector_cross(x, y, z);
    lib3ds_vector_cross(z, x, y);
    lib3ds_vector_normalize(x);
    lib3ds_vector_normalize(z);

    lib3ds_matrix_identity(M);
    M[0][0] = x[0];
    M[1][0] = x[1];
    M[2][0] = x[2];
    M[0][1] = y[0];
    M[1][1] = y[1];
    M[2][1] = y[2];
    M[0][2] = z[0];
    M[1][2] = z[1];
    M[2][2] = z[2];

    lib3ds_matrix_identity(matrix);
    lib3ds_matrix_rotate(matrix, roll, 0, 1, 0);
    lib3ds_matrix_mult(matrix, matrix, M);
    lib3ds_matrix_translate(matrix, -pos[0], -pos[1], -pos[2]);
}






