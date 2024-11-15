package skeleton

// // GetAllSkeleton ...
// func (d Data) GetAllSkeleton(ctx context.Context) error {
// 	var (
// 		rows *sqlx.Rows

// 		err error
// 	)

//  Wajib ditambahkan pointer di setiap prepare statement
// 	rows, err = (*d.stmt)[Nama_Statement].QueryxContext(ctx)
// 	for rows.Next() {
// 		if err := rows.StructScan(&evaluator); err != nil {
// 			return evaluators, errors.Wrap(err, "[DATA][GetAllDataOrderOutstanding] ")
// 		}
// 		evaluators = append(evaluators, evaluator)
// 	}
// 	defer rows.Close()

// 	return evaluators, err
// }
