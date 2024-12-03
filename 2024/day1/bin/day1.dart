import 'package:day1/day1.dart' as day1;

void main(List<String> arguments) {
  var (left, right) = day1.readFile("input.txt");

  var idCount = day1.countList<int>(right);

  int similarity = 0;
  for (var id in left) {
    // There's probably a more idiomatic way to do this but I'm not a rapper
    if (idCount[id] != null) {
      similarity += idCount[id]! * id;
    }
  }

  left.sort();
  right.sort();

  var distance = 0;
  for (var i = 0; i < left.length; i++) {
    distance += (left[i] - right[i]).abs();
  }

  print("distance: $distance");
  print("similarity: $similarity");
}
